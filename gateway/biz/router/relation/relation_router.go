package relation

import (
	"gateway/biz/handler"
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	douyin.Use(mw.JWT())
	{
		relation := douyin.Group("/relation")
		{
			relation.POST("/action/", handler.RelationAction)
			relation.GET("/follow/list/", handler.RelationList)
			relation.GET("/follower/list/", handler.FollowerList)
			relation.GET("/friend/list/", handler.FriendList)
		}
		message := douyin.Group("/message")
		{
			message.POST("/action", handler.MessageAction)
			message.GET("/chat", handler.Chat)
		}
	}
}
