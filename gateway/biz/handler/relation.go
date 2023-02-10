package handler

import (
	"context"
	"fmt"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationAction
// @Description: 关注操作
// @auth Tang-YT 2023-02-10 14:59:46
// @param ginCtx
func RelationAction(ginCtx *gin.Context) {
	uid, err := strconv.Atoi(ginCtx.Query("to_user_id"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	toUserId := int64(uid)

	act, err := strconv.Atoi(ginCtx.Query("action_type"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	actionType := int32(act)

	token := ginCtx.Query(consts.AuthorizationKey)
	var relReq service.DouyinRelationActionRequest

	relReq.ActionType = &actionType
	relReq.Token = &token
	relReq.ToUserId = &toUserId

	relationService := ginCtx.Keys[consts.RelationServiceName].(service.RelationService)
	response, err := relationService.RelationAction(context.Background(), &relReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// RelationList
// @Description: 关注列表
// @auth Tang-YT 2023-02-10 15:00:27
// @param ginCtx
func RelationList(ginCtx *gin.Context) {
	uid, err := strconv.Atoi(ginCtx.Query("user_id"))
	fmt.Println("tyt uid = ", uid)
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	userId := int64(uid)
	token := ginCtx.Query(consts.AuthorizationKey)
	//claims, _ := utils.ParseToken(token)
	var relReq service.DouyinRelationFollowListRequest

	relReq.UserId = &userId
	relReq.Token = &token

	relationService := ginCtx.Keys[consts.RelationServiceName].(service.RelationService)
	response, err := relationService.FollowList(context.Background(), &relReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// FollowerList
// @Description: 粉丝列表
// @auth sinre 2023-02-09 22:13:07
// @param ginCtx
func FollowerList(ginCtx *gin.Context) {
	var floLsReq service.DouyinRelationFollowerListRequest
	uid := ginCtx.Query("user_id")
	token := ginCtx.Query(consts.AuthorizationKey)
	uiid, err := strconv.Atoi(uid)
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	userId := int64(uiid)
	floLsReq.UserId = &userId
	floLsReq.Token = &token
	relationService := ginCtx.Keys[consts.RelationServiceName].(service.RelationService)
	response, err := relationService.FollowerList(context.Background(), &floLsReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// FriendList
// @Description: 好友列表
// @auth sinre 2023-02-09 22:13:18
// @param ginCtx
func FriendList(ginCtx *gin.Context) {
	var floLsReq service.DouyinRelationFriendListRequest
	uid := ginCtx.Query("user_id")
	token := ginCtx.Query(consts.AuthorizationKey)
	uiid, err := strconv.Atoi(uid)
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	userId := int64(uiid)
	floLsReq.UserId = &userId
	floLsReq.Token = &token
	relationService := ginCtx.Keys[consts.RelationServiceName].(service.RelationService)
	response, err := relationService.FriendList(context.Background(), &floLsReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}
