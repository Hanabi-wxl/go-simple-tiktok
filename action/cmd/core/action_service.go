package core

import (
	"action/cmd/dal/db"
	"action/cmd/model"
	"action/cmd/pack"
	"action/cmd/service"
	"action/pkg/errno"
	"action/pkg/utils"
	"context"
)

func (*ActionService) FavoriteAction(_ context.Context, req *service.DouyinFavoriteActionRequest, resp *service.DouyinFavoriteActionResponse) error {
	acType := req.GetActionType()
	token := req.GetToken()
	videoId := req.GetVideoId()
	uid := utils.GetUserId(token)
	if exit := db.CheckVideoExit(videoId); exit {
		if acType == 1 {
			db.CreateFavorite(videoId, uid)
		} else {
			db.DeleteFavorite(videoId, uid)
		}
	} else {
		return errno.VideoNotExit
	}
	pack.BuildFavoriteActionResp(resp)
	return nil
}

func (*ActionService) FavoriteList(_ context.Context, req *service.DouyinFavoriteListRequest, resp *service.DouyinFavoriteListResponse) error {
	var (
		videos     []model.Video
		videoInfos []model.VideoInfo
		videoInfo  model.VideoInfo
		author     model.Author
		userId     int64
		fids       []int64
	)
	token := req.GetToken()
	checkId := req.GetUserId()
	userId = utils.GetUserId(token)

	// 作者信息
	infoById := db.GetUserInfoById(checkId)
	author.Name = infoById.Name
	author.Id = checkId
	// 查询自己时userId == checkId
	followInfo := db.GetUserFollowInfo(checkId, userId)
	author.IsFollow = followInfo.IsFollow
	author.FollowCount = followInfo.FollowerCount
	author.FollowerCount = followInfo.FollowerCount

	// 查询视频作者喜欢列表
	favorites := db.GetFavoriteListByUserId(checkId)
	for _, favorite := range favorites {
		fids = append(fids, favorite.VideoId)
	}

	videos = db.GetFavoriteVideos(fids)

	for _, video := range videos {
		// 视频信息
		videoInfo.Id = video.VideoId
		videoInfo.Title = video.Title
		videoInfo.PlayUrl = video.PlayUrl
		videoInfo.CoverUrl = video.CoverUrl
		authId := video.Author
		// 点赞评论信息
		actionInfo := db.GetActionCount(authId)
		videoInfo.FavoriteCount = actionInfo.FavoriteCount
		videoInfo.CommentCount = actionInfo.CommentCount
		checkFavorite := db.CheckFavorite(userId, authId)
		videoInfo.IsFavorite = checkFavorite
		// 作者信息
		videoInfo.Author = author
		// 合并到全部所需信息
		videoInfos = append(videoInfos, videoInfo)
	}
	pack.BuildFavoriteListResp(resp, videoInfos)
	return nil
}

func (*ActionService) CommentAction(_ context.Context, req *service.DouyinCommentActionRequest, resp *service.DouyinCommentActionResponse) error {
	var (
		userId     int64
		videoId    = req.GetVideoId()
		acType     = req.GetActionType()
		comment    model.Comment
		userInfo   model.User
		followInfo model.FollowInfo
	)
	token := req.GetToken()
	userId = utils.GetUserId(token)
	if exit := db.CheckVideoExit(videoId); exit {
		if acType == 1 {
			comment := db.CreateComment(userId, videoId, req.GetCommentText())
			userInfo := db.GetUserInfoById(userId)
			followInfo := db.GetUserFollowInfo(userId, userId)
			pack.BuildCommentActionResp(resp, comment, userInfo, followInfo)
		} else if acType == 2 {
			db.DeleteComment(userId, videoId, req.GetCommentId())
			pack.BuildCommentActionResp(resp, comment, userInfo, followInfo)
		}
	} else {
		return errno.VideoNotExit
	}
	return nil
}

func (*ActionService) CommentList(_ context.Context, req *service.DouyinCommentListRequest, resp *service.DouyinCommentListResponse) error {
	userId := utils.GetUserId(req.GetToken())
	var (
		commentInfo  model.CommentInfo
		commentInfos []model.CommentInfo
	)
	videoId := req.GetVideoId()
	commentList := db.GetCommentList(videoId)
	for i := 0; i < len(commentList); i++ {
		comment := commentList[i]
		// 用户信息
		userInfo := db.GetUserInfoById(comment.UserId)
		// 点赞评论信息
		followInfo := db.GetUserFollowInfo(comment.UserId, userId)
		commentInfo.Id = comment.Id
		commentInfo.Content = comment.CommentText
		commentInfo.CreateDate = comment.CommentTime.Format("01-02")
		commentInfo.User.Id = userInfo.UserId
		commentInfo.User.Name = userInfo.Name
		commentInfo.User.FollowCount = followInfo.FollowCount
		commentInfo.User.FollowerCount = followInfo.FollowerCount
		commentInfo.User.IsFollow = followInfo.IsFollow
		commentInfos = append(commentInfos, commentInfo)
	}
	pack.BuildCommentListResp(resp, commentInfos)
	return nil
}
