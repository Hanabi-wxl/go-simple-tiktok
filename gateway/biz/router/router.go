package router

import (
	"gateway/biz/router/action"
	"gateway/biz/router/core"
	"gateway/biz/router/relation"
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(mw.Cors(), mw.InitMiddleware(service), mw.ErrorMiddleware())

	action.Register(ginRouter)
	core.Register(ginRouter)
	relation.Register(ginRouter)

	return ginRouter
}
