# go-simple-tiktok 简易抖音

### 项目介绍

- 视频Feed流: 支持所有用户刷抖音，视频按投稿时间倒序推出
- 视频投稿: 支持登录用户自己拍视频投稿
- 个人主页: 支持查看用户基本信息和投稿列表，注册用户流程简化
- 喜欢列表: 登录用户可以对视频点赞，在个人主页喜欢Tab下能够查看点赞视频列表
- 用户评论: 支持未登录用户查看视频下的评论列表，登录用户能够发表评论
- 关系列表: 登录用户可以关注其他用户，能够在个人主页查看本人的关注数和粉丝数，查看关注列表和粉丝列表
- 消息: 登录用户在消息页展示已关注的用户列表，点击用户头像进入聊天页后可以发送消息

### 目录结构
```
.
├─action 互动模块
│  ├─cmd 类似java src文件夹
│  │  ├─core 实现类
│  │  ├─dal 数据库操作
│  │  │  └─db 类似java mapper持久层
│  │  │  └─redis redis操作
│  │  ├─model 数据库实体对象
│  │  ├─mq 消息队列
│  │  ├─pack 封装对象
│  │  └─service 放置proto生成的代码
│  ├─idl proto文件
│  └─pkg 其它可供全局访问的代码
│      ├─consts 常量、配置
│      ├─errno 异常信息
│      └─utils 工具包
├─core 基础模块
│  ├─cmd
│  │  ├─core
│  │  ├─dal
│  │  │  └─db
│  │  ├─model
│  │  ├─pack
│  │  └─service
│  ├─idl
│  └─pkg
│      ├─consts
│      ├─errno
│      └─utils
├─gateway 网关 （gin）
│  ├─lib 二进制可执行文件
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
│  │  ├─consts 常量、配置
│  │  ├─doc 文本文件
│  │  ├─result 结果类、异常类
│  │  ├─sql 数据库结构文件
│  │  └─utils 工具包
│  └─static 静态资源文件
└─relation 社交模块
  ├─cmd
  │  ├─core
  │  ├─dal
  │  │  └─db
  │  ├─model
  │  ├─pkg
  │  └─service
  ├─idl
  └─pkg
     ├─consts
     ├─errno
     └─utils
```

### 项目架构图

![image-20230210021530831](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/image-20230210021530831.png)

### 技术点

- web: gin
- rpc: go-micro(v2), protobuf, etcd
- 数据库: mysql
- 持久层: gorm
- 认证: JWT
- 参数校验: protoc-gen-validate
- 视频文件操作: ffmpeg


### 数据库结构

![img](https://sinre.oss-cn-beijing.aliyuncs.com/picgo/img.png)


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
