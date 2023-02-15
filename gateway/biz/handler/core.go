package handler

import (
	"context"
	"encoding/json"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

// Feed
// @Description: 视频流接口
// @auth sinre 2023-02-07 16:46:29
// @param ginCtx
func Feed(ginCtx *gin.Context) {
	var feedReq service.DouyinFeedRequest
	var latestTime int64
	latestTimeStr := ginCtx.Query("latest_time")
	if len(latestTimeStr) != len(consts.DefaultTime) {
		latestTime = time.Now().UnixMilli()
	} else {
		atoi, _ := strconv.Atoi(latestTimeStr)
		latestTime = int64(atoi)
	}
	token := ginCtx.Query(consts.AuthorizationKey)
	feedReq.LatestTime = &latestTime
	feedReq.Token = &token
	// 获取core服务
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	// 调用服务
	response, err := coreService.Feed(ginCtx, &feedReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	// 返回结果
	ginCtx.JSON(http.StatusOK, response)
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
	regReq.Password = &password
	// 数据校验
	if err := regReq.Validate(); err != nil {
		SendValidateErr(ginCtx, err)
		return
	}
	regReq.Username = &username
	regReq.Password = &password

	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.UserRegister(context.Background(), &regReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	// 颁发token
	token, _ := utils.GenerateToken(response.GetUserId())
	response.Token = &token
	ginCtx.JSON(http.StatusOK, response)
}

// Login
// @Description: 用户登录
// @auth sinre 2023-02-07 16:46:53
// @param ginCtx
func Login(ginCtx *gin.Context) {
	var loginReq service.DouyinUserLoginRequest
	username := ginCtx.Query("username")
	password := ginCtx.Query("password")
	loginReq.Username = &username
	loginReq.Password = &password
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.UserLogin(context.Background(), &loginReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	token, _ := utils.GenerateToken(response.GetUserId())
	response.Token = &token
	ginCtx.JSON(http.StatusOK, response)
}

// User
// @Description: 用户信息
// @auth sinre 2023-02-07 18:08:08
// @param ginCtx
func User(ginCtx *gin.Context) {
	var userReq service.DouyinUserRequest
	uid, err := strconv.Atoi(ginCtx.Query("user_id"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	userId := int64(uid)
	token := ginCtx.Query(consts.AuthorizationKey)
	userReq.UserId = &userId
	userReq.Token = &token
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.User(context.Background(), &userReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// PublishAction
// @Description: 投稿接口
// @auth sinre 2023-02-07 16:47:18
// @param ginCtx
func PublishAction(ginCtx *gin.Context) {
	title := ginCtx.PostForm("title")
	title = utils.Filter.Replace(title, '~')
	token := ginCtx.PostForm(consts.AuthorizationKey)
	file, err := ginCtx.FormFile("data")
	if file == nil {
		SendClientErr(ginCtx, consts.FileNotFoundErr)
		return
	}
	if file.Size > consts.MaxFileSize {
		SendClientErr(ginCtx, consts.FileToLargeErr)
		return
	}
	if err != nil {
		SendClientErr(ginCtx, consts.PostFormVideoErr)
		return
	}
	suffix := filepath.Ext(file.Filename)
	//判断是否为视频格式
	if _, ok := consts.VideoTypeMap[suffix]; !ok {
		SendClientErr(ginCtx, consts.VideoTypeErrErr)
		return
	}
	// 使用uuid作为文件名
	name := utils.NewFileName()
	filename := name + suffix
	savePath := filepath.Join("./static", filename)
	err = ginCtx.SaveUploadedFile(file, savePath)
	if err != nil {
		SendClientErr(ginCtx, consts.SaveFileTempErr)
		return
	}
	// 截取封面
	err = utils.Capture(name, false)
	if err != nil {
		SendClientErr(ginCtx, consts.VideoCaptureErr)
		return
	}
	claims, _ := utils.ParseToken(token)
	mapData := map[string]interface{}{
		"Author": map[string]int64{
			"Id": claims.UserId,
		},
		"play_url":  utils.GetFileUrl(filename),
		"cover_url": utils.GetFileUrl(name + ".jpg"),
	}

	bytes, _ := json.Marshal(mapData)
	var pubReq service.DouyinPublishActionRequest
	pubReq.Title = &title
	pubReq.Token = &token
	pubReq.Data = bytes

	// 调用服务将文件信息保存至数据库
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.PublishAction(context.Background(), &pubReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// PublishList
// @Description: 发布列表
// @auth sinre 2023-02-07 16:47:25
// @param ginCtx
func PublishList(ginCtx *gin.Context) {
	var pubReq service.DouyinPublishListRequest
	token := ginCtx.Query(consts.AuthorizationKey)
	uid, err := strconv.Atoi(ginCtx.Query("user_id"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	userId := int64(uid)
	pubReq.UserId = &userId
	pubReq.Token = &token
	coreService := ginCtx.Keys[consts.CoreServiceName].(service.CoreService)
	response, err := coreService.PublishList(context.Background(), &pubReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}
