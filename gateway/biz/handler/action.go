package handler

import (
	"context"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction
// @Description: 赞操作
// @auth sinre 2023-02-09 16:27:39
// @param ginCtx
func FavoriteAction(ginCtx *gin.Context) {
	var (
		favAcReq service.DouyinFavoriteActionRequest
		vid      int64
		tid      int32
	)
	token := ginCtx.Query(consts.AuthorizationKey)
	vvid, err1 := strconv.Atoi(ginCtx.Query("video_id"))
	ttid, err2 := strconv.Atoi(ginCtx.Query("action_type"))
	if err1 != nil || err2 != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	vid = int64(vvid)
	tid = int32(ttid)
	favAcReq.Token = &token
	favAcReq.ActionType = &tid
	favAcReq.VideoId = &vid
	actionService := ginCtx.Keys[consts.ActionServiceName].(service.ActionService)
	response, err := actionService.FavoriteAction(context.Background(), &favAcReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// FavoriteList
// @Description: 喜欢列表
// @auth sinre 2023-02-09 16:27:56
// @param ginCtx
func FavoriteList(ginCtx *gin.Context) {
	var (
		favLsReq service.DouyinFavoriteListRequest
		userId   int64
	)
	token := ginCtx.Query(consts.AuthorizationKey)
	uid, err := strconv.Atoi(ginCtx.Query("user_id"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	userId = int64(uid)
	favLsReq.Token = &token
	favLsReq.UserId = &userId
	actionService := ginCtx.Keys[consts.ActionServiceName].(service.ActionService)
	response, err := actionService.FavoriteList(context.Background(), &favLsReq)
	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// CommentAction
// @Description: 评论操作
// @auth sinre 2023-02-09 16:28:07
// @param ginCtx
func CommentAction(ginCtx *gin.Context) {
	var (
		favLsReq  service.DouyinCommentActionRequest
		videoId   int64
		acType    int32
		commentId int64
		cid       int64
	)
	token := ginCtx.Query(consts.AuthorizationKey)
	commentText := ginCtx.Query("comment_text")
	vid, err1 := strconv.Atoi(ginCtx.Query("video_id"))
	cidStr := ginCtx.Query("comment_id")
	if len(cidStr) != 0 {
		ccid, err2 := strconv.Atoi(cidStr)
		if err2 != nil {
			SendClientErr(ginCtx, consts.ParamErr)
			return
		}
		cid = int64(ccid)
	}
	tt, err3 := strconv.Atoi(ginCtx.Query("action_type"))
	if err1 != nil || err3 != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	videoId = int64(vid)
	commentId = int64(cid)
	acType = int32(tt)

	favLsReq.Token = &token
	favLsReq.VideoId = &videoId
	favLsReq.CommentId = &commentId
	favLsReq.ActionType = &acType
	favLsReq.CommentText = &commentText

	actionService := ginCtx.Keys[consts.ActionServiceName].(service.ActionService)
	response, err := actionService.CommentAction(context.Background(), &favLsReq)

	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}

// CommentList
// @Description: 评论列表
// @auth sinre 2023-02-09 16:28:17
// @param ginCtx
func CommentList(ginCtx *gin.Context) {
	var (
		coLsReq service.DouyinCommentListRequest
		videoId int64
	)
	token := ginCtx.Query(consts.AuthorizationKey)
	vid, err := strconv.Atoi(ginCtx.Query("video_id"))
	if err != nil {
		SendClientErr(ginCtx, consts.ParamErr)
		return
	}
	videoId = int64(vid)
	coLsReq.Token = &token
	coLsReq.VideoId = &videoId
	actionService := ginCtx.Keys[consts.ActionServiceName].(service.ActionService)
	response, err := actionService.CommentList(context.Background(), &coLsReq)

	if err != nil {
		SendServiceErr(ginCtx, err)
		return
	}
	ginCtx.JSON(http.StatusOK, response)
}
