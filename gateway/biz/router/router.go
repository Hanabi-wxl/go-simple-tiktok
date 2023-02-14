package router

import (
	"gateway/biz/router/action"
	"gateway/biz/router/core"
	"gateway/biz/router/relation"
	"gateway/mw"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(mw.Cors(), mw.InitMiddleware(service), mw.ErrorMiddleware())

	// 开启静态资源文件夹
	ginRouter.StaticFS("/static", http.Dir("./static"))

	// 注册路由
	action.Register(ginRouter)
	core.Register(ginRouter)
	relation.Register(ginRouter)
	return ginRouter
}
