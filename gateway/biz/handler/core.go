package handler

import (
	"context"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"github.com/gin-gonic/gin"
)

func Feed(ginCtx *gin.Context) {
	var feedReq service.DouyinFeedRequest
	// 参数绑定
	err := ginCtx.Bind(&feedReq)
	if err != nil {
		// 异常结果返回
		SendResponse(ginCtx, consts.ParamErr, nil)
	}
	// 获取core服务
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	// 调用服务
	response, err := coreService.Feed(context.Background(), &feedReq)
	// 返回结果
	ginCtx.JSON(consts.SuccessCode, response)
}
