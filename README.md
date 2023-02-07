# go-simple-tiktok

###

目录结构
```
.
├─action 互动模块
│  ├─cmd 类似java src文件夹
│  │  ├─core 实现类
│  │  ├─dal 数据库操作
│  │  │  └─db 类似java mapper持久层
│  │  ├─pack 封装对象
│  │  └─service 放置proto生成的代码
│  ├─idl proto文件
│  └─pkg 其它可供全局访问的代码
│      └─consts 常量、配置
├─core 基础模块
│  ├─cmd
│  │  ├─core
│  │  ├─dal
│  │  │  └─db
│  │  ├─pack
│  │  └─service
│  ├─idl
│  └─pkg
│      └─consts
├─gateway 网关 （gin）
│  ├─biz  类似java src文件夹
│  │  ├─handler 接口处理函数
│  │  ├─router 总路由
│  │  │  ├─action 路由、中间件
│  │  │  ├─core 路由、中间件
│  │  │  └─relation 路由、中间件
│  │  ├─service 放置proto生成的代码
│  │  └─wrappers 装饰器 可配置服务熔断
│  ├─idl proto文件
│  ├─mw 全局中间件
│  └─pkg 其它可供全局访问的代码
│      ├─consts 常量、配置
│      ├─result 结果类、异常类
│      └─utils 工具包
└─relation 社交模块
├─cmd
│  ├─core
│  ├─dal
│  │  └─db
│  ├─pkg
│  └─service
├─idl
└─pkg
└─consts
```

### Sample

1. 启动etcd
2. debug模式启动gateway服务
3. debug模式启动core服务
4. 查看路由接口的定义函数

``` go
// 文件目录：gateway/biz/router/core/core_router.go
func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	{
		douyin.GET("/feed", handler.Feed)
	}
}
```

5. 查看handler处理函数

``` go
// 文件目录：gateway/biz/handler/core.go
func Feed(ginCtx *gin.Context) {
	var feedReq service.DouyinFeedRequest
	// 参数绑定
	err := ginCtx.Bind(&feedReq)
	if err != nil {
		// 异常结果返回
		SendResponse(ginCtx, consts.ParamErr, nil)
	}
	// 获取core服务
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	// 调用服务
	response, err := coreService.Feed(context.Background(), &feedReq)
	// 返回结果
	ginCtx.JSON(consts.SuccessCode, response)
}
```

6. gataway服务加断点

![image-20230207114623319](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/image-20230207114623319.png)

7. 访问127.0.0.1:8000/douyin/feed
8. 第一处断点

![image-20230207114744023](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/image-20230207114744023.png)

9. 第二处断点 准备进入core服务

![image-20230207114800820](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/image-20230207114800820.png)

9. 查看core服务处理函数

``` go
// 文件目录：core/cmd/core/core_service.go
func (*CoreService) Feed(ctx context.Context, req *service.DouyinFeedRequest, resp *service.DouyinFeedResponse) error {
	fmt.Println("qqqq")
	return nil
}
```

10. 加断点

![image-20230207114200465](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/image-20230207114200465.png)

11. 执行微服务代码逻辑

![image-20230207114247967](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/image-20230207114247967.png)

