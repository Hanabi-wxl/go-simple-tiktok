package main

import (
	ccore "core/cmd/core"
	"core/cmd/dal"
	"core/cmd/service"
	"core/pkg/consts"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	dal.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(consts.ETCDAddr),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(consts.CoreServiceName), // 微服务名字
		micro.Address(consts.CoreServiceAddr),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = service.RegisterCoreServiceHandler(microService.Server(), new(ccore.CoreService))
	// 启动微服务
	_ = microService.Run()
}
