package handler

import (
	"gateway/biz/service"
	"gateway/pkg/consts"
	"github.com/gin-gonic/gin"
)

func Feed(ginCtx *gin.Context) {
	var feedReq service.DouyinFeedRequest
	err := ginCtx.Bind(&feedReq)
	if err != nil {
		SendResponse(ginCtx, consts.ParamErr, nil)
	}
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.Feed(ginCtx, &feedReq)
	ginCtx.JSON(consts.SuccessCode, response)
}
