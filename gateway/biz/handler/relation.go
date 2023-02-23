package handler

import (
	"context"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	if err != nil || (act != 1 && act != 2) {
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
	if response.UserList == nil {
		SendMap(ginCtx, response, "user_list")
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
	if response.UserList == nil {
		SendMap(ginCtx, response, "user_list")
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// MessageAction
// @Description: 发送消息
// @auth since 2023-02-11 15:37:52
// @param ginCtx
func MessageAction(ginCtx *gin.Context) {
	var msgAct service.DouyinMessageActionRequest
	toId, err := strconv.Atoi(ginCtx.Query("to_user_id"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	toUserId := int64(toId)
	content := ginCtx.Query("content")
	if len(content) == 0 {
		SendClientErr(ginCtx, consts.NoContentErr)
		return
	}
	token := ginCtx.Query(consts.AuthorizationKey)
	act, err := strconv.Atoi(ginCtx.Query("action_type"))
	if err != nil || (act != 1) {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	actionType := int32(act)

	msgAct.ToUserId = &toUserId
	msgAct.Token = &token
	msgAct.ActionType = &actionType
	msgAct.Content = &content

	relationService := ginCtx.Keys[consts.RelationServiceName].(service.RelationService)
	ctx := context.Background()
	response, err := relationService.MessageAction(ctx, &msgAct)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)

}

// Chat
// @Description: 聊天记录
// @auth sinre 2023-02-13 11:46:27
// @param ginCtx
func Chat(ginCtx *gin.Context) {
	var msctReq service.DouyinMessageChatRequest
	tuid, err := strconv.Atoi(ginCtx.Query("to_user_id"))
	token := ginCtx.Query(consts.AuthorizationKey)
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	ttuid := int64(tuid)
	msctReq.ToUserId = &ttuid
	msctReq.Token = &token
	relationService := ginCtx.Keys[consts.RelationServiceName].(service.RelationService)
	response, err := relationService.MessageChat(context.Background(), &msctReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	if response.MessageList == nil {
		SendMap(ginCtx, response, "message_list")
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}
