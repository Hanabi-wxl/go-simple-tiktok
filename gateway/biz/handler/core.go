package handler

import (
	"context"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"github.com/gin-gonic/gin"
)

// Feed
// @Description: 视频流接口
// @auth sinre 2023-02-07 16:46:29
// @param ginCtx
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

// Register
// @Description: 用户注册
// @auth sinre 2023-02-07 16:46:44
// @param ginCtx
func Register(ginCtx *gin.Context) {
	var regReq service.DouyinUserRegisterRequest
	username := ginCtx.Query("username")
	password := ginCtx.Query("password")
	regReq.Username = &username
	regReq.Username = &password
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.UserRegister(context.Background(), &regReq)
	if err != nil {
		SendResponse(ginCtx, consts.ParamErr, nil)
	}
	ginCtx.JSON(consts.SuccessCode, response)
}

// Login
// @Description: 用户登录
// @auth sinre 2023-02-07 16:46:53
// @param ginCtx
func Login(ginCtx *gin.Context) {

}

// User
// @Description: 用户信息
// @auth sinre 2023-02-07 18:08:08
// @param ginCtx
func User(ginCtx *gin.Context) {

}

// PublishAction
// @Description: 投稿接口
// @auth sinre 2023-02-07 16:47:18
// @param ginCtx
func PublishAction(ginCtx *gin.Context) {

}

// PublishList
// @Description: 发布列表
// @auth sinre 2023-02-07 16:47:25
// @param ginCtx
func PublishList(ginCtx *gin.Context) {

}
