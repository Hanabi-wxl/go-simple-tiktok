package core

import (
	"gateway/biz/handler"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	{
		douyin.GET("/feed", handler.Feed)

		user := douyin.Group("/user")
		{
			user.POST("/register", handler.Register)
			user.POST("/login", handler.Login)
			user.GET("/", handler.User)
		}

		publish := douyin.Group("/publish")
		{
			publish.POST("/action", handler.PublishAction)
			publish.GET("/list", handler.PublishList)
		}
	}
}
