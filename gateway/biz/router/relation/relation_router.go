package relation

import (
	"gateway/biz/handler"
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	{
		relation := douyin.Group("/relation")
		{
			relation.POST("/action/", mw.JWT(), handler.RelationAction)
			relation.GET("/follow/list/", handler.RelationList)
			relation.GET("/follower/list/", handler.FollowerList)
			relation.GET("/friend/list/", mw.JWT(), handler.FriendList)
		}
		message := douyin.Group("/message")
		message.Use(mw.JWT())
		{
			message.POST("/action/", handler.MessageAction)
			message.GET("/chat/", handler.Chat)
		}
	}
}
