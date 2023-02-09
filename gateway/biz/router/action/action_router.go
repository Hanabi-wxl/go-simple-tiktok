package action

import (
	"gateway/biz/handler"
	"gateway/mw"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	{
		favorite := douyin.Group("/favorite")
		{
			favorite.POST("/action/", mw.JWT(), handler.FavoriteAction)
			favorite.GET("/list/", handler.FavoriteList)
		}
		comment := douyin.Group("/comment")
		{
			comment.POST("/action/", mw.JWT(), handler.CommentAction)
			comment.GET("/list/", handler.CommentList)
		}
	}
}
