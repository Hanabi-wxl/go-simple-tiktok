package mw

import (
	"encoding/json"
	"gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

// ErrorMiddleware 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				var serviceError result.ServiceError
				_ = json.Unmarshal([]byte(err.Error()), &serviceError)
				context.JSON(int(serviceError.Code), gin.H{
					"status_code": serviceError.Code,
					"status_msg":  serviceError.Detail,
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
