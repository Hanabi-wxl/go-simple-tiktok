package handler

import (
	"gateway/pkg/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendServiceErr 服务端异常响应
func SendServiceErr(c *gin.Context, err error) {
	serviceError := result.ConvertServiceErr(err.Error())
	c.JSON(http.StatusBadRequest, gin.H{
		"status_code": serviceError.Code,
		"status_msg":  serviceError.Detail,
	})
}

// SendClientErr 客户端异常响应
func SendClientErr(c *gin.Context, err result.ClientError) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status_code": err.StatusCode,
		"status_msg":  err.StatusMsg,
	})
}
