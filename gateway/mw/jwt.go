package mw

import (
	"gateway/pkg/consts"
	"gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		var code int
		token := ginCtx.Query(consts.AuthorizationKey)
		if token == "" {
			token = ginCtx.PostForm(consts.AuthorizationKey)
			// 无token异常
			if token == "" {
				code = consts.NoTokenErrCode
			}
		}
		if code == 0 {
			_, err := utils.ParseToken(token)
			// 解析token异常
			if err != nil {
				code = consts.AuthorizationErrCode
			}
		}
		if code != 0 {
			ginCtx.JSON(code, gin.H{
				"status_code": code,
				"status_msg":  "鉴权失败",
			})
			ginCtx.Abort()
		}
		ginCtx.Next()
	}
}
