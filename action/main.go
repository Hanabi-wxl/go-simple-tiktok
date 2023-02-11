package main

import (
	"action/cmd/core"
	"action/cmd/dal"
	"action/cmd/mq"
	"action/cmd/service"
	"action/pkg/consts"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func Init() {
	dal.Init()
	mq.Init()
}

func main() {
	Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(consts.ETCDAddr),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(consts.ActionServiceName), // 微服务名字
		micro.Address(consts.ActionServiceAddr),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = service.RegisterActionServiceHandler(microService.Server(), new(core.ActionService))
	// 启动微服务
	_ = microService.Run()
}
