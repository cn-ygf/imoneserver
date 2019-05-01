package service

import (
	"github.com/cn-ygf/imoneserver"
	"sync"
	"sync/atomic"
)

type ServiceManager interface {
	Add(imoneserver.Service)               // 添加服务
	Remove(string)                         // 移除服务
	GetService(string) imoneserver.Service // 根据别名获取服务
	Count() int64                          // 获取服务总数
	CloseAllService()                      // 关闭所有服务
}

// 服务管理结构
type CoreServiceManager struct {
	services sync.Map // 保存服务map
	count    int64    // 服务计数
	startId  int64    // 起始服务id
}

func (core *CoreServiceManager) Add(service imoneserver.Service) {
	atomic.AddInt64(&core.count, 1)
	if _, ok := service.(imoneserver.ServiceProperty); ok {
		core.services.Store(service.(imoneserver.ServiceProperty).Name(), service)
	} else {
		core.services.Store(service.TypeName(), service)
	}

}

func (core *CoreServiceManager) Remove(name string) {
	core.services.Delete(name)
	atomic.AddInt64(&core.count, -1)
}

func (core *CoreServiceManager) GetService(name string) imoneserver.Service {
	if v, ok := core.services.Load(name); ok {
		return v.(imoneserver.Service)
	}
	return nil
}

func (core *CoreServiceManager) Count() int64 {
	return atomic.LoadInt64(&core.count)
}

func (core *CoreServiceManager) CloseAllService() {
	core.VisitTunnel(func(service imoneserver.Service) bool {
		service.Close()
		return true
	})
}

// 遍历所有通道
func (core *CoreServiceManager) VisitTunnel(callback func(imoneserver.Service) bool) {
	core.services.Range(func(key, value interface{}) bool {
		return callback(value.(imoneserver.Service))
	})
}
