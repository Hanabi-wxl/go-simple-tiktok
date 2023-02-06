package main

import (
	"gateway/biz/router"
	"gateway/biz/service"
	"gateway/biz/wrappers"
	"gateway/pkg/consts"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(consts.ETCDAddr),
	)
	coreMicroService := micro.NewService(
		micro.Name(consts.CoreClientName),
		micro.WrapClient(wrappers.NewCoreWrapper),
	)
	// 服务调用实例
	coreService := service.NewCoreService(consts.CoreServiceName, coreMicroService.Client())

	actionMicroService := micro.NewService(
		micro.Name(consts.ActionClientName),
		micro.WrapClient(wrappers.NewActionWrapper),
	)
	actionService := service.NewActionService(consts.ActionServiceName, actionMicroService.Client())

	relationMicroService := micro.NewService(
		micro.Name(consts.RelationClientName),
		micro.WrapClient(wrappers.NewActionWrapper),
	)
	relationService := service.NewActionService(consts.RelationServiceName, relationMicroService.Client())

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name(consts.GateWayServiceName),
		web.Address(consts.GateWayAddr),
		web.Handler(router.NewRouter(coreService, actionService, relationService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	_ = server.Init()
	_ = server.Run()
}
