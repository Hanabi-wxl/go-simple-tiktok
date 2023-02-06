package mw

import (
	"gateway/pkg/consts"
	"github.com/gin-gonic/gin"
)

// InitMiddleware 接受服务实例，并存到gin.Key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存在gin.Keys中
		context.Keys = make(map[string]interface{})
		context.Keys[consts.CoreServiceName] = service[0]
		context.Keys[consts.ActionServiceName] = service[1]
		context.Keys[consts.RelationServiceName] = service[2]
		context.Next()
	}
}
