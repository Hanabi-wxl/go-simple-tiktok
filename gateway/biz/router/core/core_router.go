package core

import (
	"gateway/biz/handler"
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	{
		douyin.GET("/feed/", handler.Feed)

		user := douyin.Group("/user")
		{
			user.POST("/register/", handler.Register)
			user.POST("/login/", handler.Login)
			user.GET("/", mw.JWT(), handler.User)
		}

		publish := douyin.Group("/publish")
		{
			publish.POST("/action/", mw.JWT(), handler.PublishAction)
			publish.GET("/list/", handler.PublishList)
		}

		relation := douyin.Group("/relation")
		//relation.Use(mw.JWT())
		{
			relation.POST("/action", handler.RelationAction)
			relation.GET("/follow/list", handler.RelationList)
		}
	}
}
