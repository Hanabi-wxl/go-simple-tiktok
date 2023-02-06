package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"relation/cmd/core"
	"relation/cmd/service"
	"relation/pkg/consts"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(consts.ETCDAddr),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(consts.RelationServiceName), // 微服务名字
		micro.Address(consts.RelationServiceAddr),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = service.RegisterRelationServiceHandler(microService.Server(), new(core.RelationService))
	// 启动微服务
	_ = microService.Run()
}
