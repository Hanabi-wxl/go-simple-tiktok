package relation

import (
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	douyin.Use(mw.JWT())
	{
		relation := douyin.Group("/relation")
		{
			relation.POST("/action/")
			relation.GET("/follow/list/")
			relation.GET("/follower/list/")
			relation.GET("/friend/list/")
		}
		message := douyin.Group("/message")
		{
			message.POST("/action")
			message.GET("/chat")
		}
	}
}
