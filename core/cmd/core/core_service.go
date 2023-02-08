package core

import (
	"context"
	"core/cmd/dal/db"
	"core/cmd/model"
	"core/cmd/pack"
	"core/cmd/service"
	"core/pkg/errno"
	"core/pkg/utils"
	"encoding/json"
	"time"
)

func (*CoreService) Feed(ctx context.Context, req *service.DouyinFeedRequest, resp *service.DouyinFeedResponse) error {
	var (
		videos     []model.Video
		videoInfos []model.VideoInfo
		videoInfo  model.VideoInfo
		author     model.Author
		lastTime   int64
	)
	// 时间戳
	lastTime = req.GetLatestTime()
	// 转换为时间对象
	timeTime := time.UnixMilli(lastTime)
	token := req.GetToken()
	claims, _ := utils.ParseToken(token)
	userId := claims.UserId
	// 获取视频列表及时间戳
	videos, lastTime = db.FeedVideos(timeTime)
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
		infoById := db.GetUserInfoById(authId)
		author.Name = infoById.Name
		author.Id = authId
		followInfo := db.GetUserFollowInfo(authId, userId)
		author.IsFollow = followInfo.IsFollow
		author.FollowCount = followInfo.FollowerCount
		author.FollowerCount = followInfo.FollowerCount
		videoInfo.Author = author
		// 合并到全部所需信息
		videoInfos = append(videoInfos, videoInfo)
	}
	pack.BuildFeedRes(resp, videoInfos, lastTime)
	return nil
}

func (*CoreService) UserRegister(ctx context.Context, req *service.DouyinUserRegisterRequest, resp *service.DouyinUserRegisterResponse) error {
	username := req.GetUsername()
	password := req.GetPassword()
	if exit := db.CheckUserExit(username); exit {
		return errno.UserAlreadyExitErr
	} else {
		db.CreateUser(username, password)
		user := db.GetUserInfoByUsername(username)
		pack.BuildUserRegResp(resp, &user)
	}
	return nil
}

func (*CoreService) UserLogin(ctx context.Context, req *service.DouyinUserLoginRequest, resp *service.DouyinUserLoginResponse) error {
	var user model.User
	user.Name = req.GetUsername()
	if exit := db.CheckUserExit(user.Name); !exit {
		return errno.UserNotExitErr
	}
	userInfo := db.GetUserInfoByUsername(user.Name)
	// 设置为数据库内的加密密码
	user.Password = userInfo.Password
	// 检查输入的密码是否争取
	if check := user.CheckPassword(req.GetPassword()); !check {
		return errno.PasswordIncorrectErr
	}
	pack.BuildUserLoginResp(resp, &userInfo)
	return nil
}

func (*CoreService) User(ctx context.Context, req *service.DouyinUserRequest, resp *service.DouyinUserResponse) error {
	// 查看自己信息时checkUserId == userId
	// 查看他人信息时userId为抖音使用者id checkUserId为发视频作者id
	checkUserId := req.GetUserId()
	claims, _ := utils.ParseToken(req.GetToken())
	userId := claims.UserId
	// 获取用户信息
	checkUserInfo := db.GetUserInfoById(checkUserId)
	followInfo := db.GetUserFollowInfo(checkUserId, userId)
	pack.BuildUserResp(resp, &checkUserInfo, followInfo)
	return nil
}

func (*CoreService) PublishAction(ctx context.Context, req *service.DouyinPublishActionRequest, resp *service.DouyinPublishActionResponse) error {
	var fileInfo service.Video
	var video model.Video
	_ = json.Unmarshal(req.Data, &fileInfo)
	video.Author = fileInfo.GetAuthor().GetId()
	video.PlayUrl = fileInfo.GetPlayUrl()
	video.CoverUrl = fileInfo.GetCoverUrl()
	video.Title = req.GetTitle()
	db.CreateFileInfo(video)
	pack.BuildPublishActionResp(resp)
	return nil
}

func (*CoreService) PublishList(ctx context.Context, req *service.DouyinPublishListRequest, resp *service.DouyinPublishListResponse) error {
	return nil
}
