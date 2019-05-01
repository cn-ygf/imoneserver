package service

import (
	"fmt"
	"github.com/cn-ygf/imoneserver"
)

var serviceMgr ServiceManager

func init() {
	if serviceMgr == nil {
		serviceMgr = &CoreServiceManager{}
	}
}

type ServiceCreateFunc func() imoneserver.Service

var serviceByName = map[string]ServiceCreateFunc{}

// 注册Service创建器
func RegisterServiceCreator(f ServiceCreateFunc) {
	// 临时实例化一个，获取类型
	dummyService := f()
	if _, ok := serviceByName[dummyService.TypeName()]; ok {
		panic("duplicate peer type: " + dummyService.TypeName())
	}
	serviceByName[dummyService.TypeName()] = f
}

// 创建一个Service
func NewService(serviceType ...string) imoneserver.Service {
	if len(serviceType) < 1 {
		panic(fmt.Sprintf("service type not found '%s'\ntry to add code below:\nimport (\n  _ \"%s\"\n)\n\n",
			serviceType,
			"TODO"))
	}
	serviceCreator := serviceByName[serviceType[0]]
	if serviceCreator == nil {
		panic(fmt.Sprintf("service type not found '%s'\ntry to add code below:\nimport (\n  _ \"%s\"\n)\n\n",
			serviceType,
			"TODO"))
	}

	s := serviceCreator()
	if len(serviceType) > 1 {
		s.(imoneserver.ServiceProperty).SetName(serviceType[1])
	}
	serviceMgr.Add(s)
	return s
}

// 获取一个Service
func GetService(name string)imoneserver.Service{
	return serviceMgr.GetService(name)
}