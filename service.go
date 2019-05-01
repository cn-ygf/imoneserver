package imoneserver

// 服务接口
type Service interface {
	Run(...interface{}) // 启动服务
	Close()             // 关闭服务
	TypeName() string   // 取得服务名
}

// 服务属性
type ServiceProperty interface {
	Name() string   // 服务别名
	SetName(string) // 设置服务别名
}
