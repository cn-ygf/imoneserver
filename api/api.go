package api

import (
	"github.com/cn-ygf/imoneserver"
	"github.com/cn-ygf/yin"
)

// api服务
type apiService struct {
	imoneserver.ServiceProperty
	http yin.Engine // web引擎
	name string     // 别名
}

func (api *apiService) Run(param ...interface{}) {
	if len(param) < 1 {
		api.http = yin.Default()
	} else {
		api.http = yin.New(param[0].(string))
	}
	RegisterRouter(api.http)
	api.http.Run()
	log.Infof("%s: service is running\n", api.Name())
}

func (api *apiService) Close() {
	api.http.Close()
	log.Infof("%s: service is closed\n", api.Name())
}

func (api *apiService) TypeName() string {
	return "api"
}

func (api *apiService) Name() string {
	if len(api.name) < 1 {
		return api.TypeName()
	}
	return api.name
}

func (api *apiService) SetName(name string) {
	api.name = name
}
