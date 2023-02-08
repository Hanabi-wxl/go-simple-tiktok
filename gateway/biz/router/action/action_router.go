package action

import (
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	douyin.Use(mw.JWT())
	{
		favorite := douyin.Group("/favorite")
		{
			favorite.POST("/action/")
			favorite.GET("/list/")
		}
		comment := douyin.Group("/comment")
		{
			comment.POST("/action/")
			comment.GET("/list/")
		}
	}
}
