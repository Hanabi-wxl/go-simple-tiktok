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
				var globalError result.GlobalError
				_ = json.Unmarshal([]byte(err.Error()), &globalError)
				context.JSON(int(globalError.Code), gin.H{
					"code": globalError.Code,
					"msg":  globalError.Detail,
					"data": nil,
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
