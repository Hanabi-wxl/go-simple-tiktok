package core

import (
	"gateway/biz/handler"
	"github.com/gin-gonic/gin"
)

func Register(engin *gin.Engine) {
	douyin := engin.Group("/douyin")
	{
		douyin.GET("/feed", handler.Feed)
	}
}
