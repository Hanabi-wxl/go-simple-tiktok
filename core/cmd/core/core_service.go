package core

import (
	"context"
	"core/cmd/dal/db"
	"core/cmd/model"
	"core/cmd/pack"
	"core/cmd/service"
	"core/pkg/consts"
	"core/pkg/errno"
	"core/pkg/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func (*CoreService) Feed(_ context.Context, req *service.DouyinFeedRequest, resp *service.DouyinFeedResponse) error {
	var (
		videos     []model.Video
		videoInfos []*service.Video
		lastTime   int64
		userId     int64
	)
	// 时间戳
	lastTime = req.GetLatestTime()
	// 转换为时间对象
	timeTime := time.UnixMilli(lastTime)
	token := req.GetToken()
	claims, _ := utils.ParseToken(token)
	// 登录后获取userId
	if claims != nil {
		userId = claims.UserId
	}
	// 获取视频列表及时间戳
	videos, lastTime = db.FeedVideos(timeTime)
	for i := 0; i < len(videos); i++ {
		var author service.User
		// 视频信息
		var videoInfo service.Video
		videoInfo.Id = &videos[i].VideoId
		videoInfo.Title = &videos[i].Title
		videoInfo.PlayUrl = &videos[i].PlayUrl
		videoInfo.CoverUrl = &videos[i].CoverUrl
		authId := videos[i].Author
		// 点赞评论信息
		actionInfo := db.GetActionCount(videos[i].VideoId)
		videoInfo.FavoriteCount = &actionInfo.FavoriteCount
		videoInfo.CommentCount = &actionInfo.CommentCount
		checkFavorite := db.CheckFavorite(userId, videos[i].VideoId)
		videoInfo.IsFavorite = &checkFavorite
		// 作者信息
		infoById := db.GetUserInfoById(authId)
		author.Name = &infoById.Name
		author.Id = &authId
		author.Avatar = &infoById.Avatar
		url := consts.BackgroundImgUrl
		author.BackgroundImage = &url
		followInfo := db.GetUserFollowInfo(authId, userId)
		author.IsFollow = &followInfo.IsFollow
		author.FollowCount = &followInfo.FollowCount
		author.FollowerCount = &followInfo.FollowerCount
		videoInfo.Author = &author
		// 合并到全部所需信息
		videoInfos = append(videoInfos, &videoInfo)
	}
	pack.BuildFeedResp(resp, videoInfos, lastTime)
	return nil
}

func (*CoreService) UserRegister(_ context.Context, req *service.DouyinUserRegisterRequest, resp *service.DouyinUserRegisterResponse) error {
	username := req.GetUsername()
	password := req.GetPassword()
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10) + 1
	avatar := fmt.Sprintf("%s%d%s", consts.AvatarFileUrl, i, ".png")
	if exist := db.CheckUserExist(username); exist {
		return errno.UserAlreadyExistErr
	} else {
		db.CreateUser(username, password, avatar)
		user := db.GetUserInfoByUsername(username)
		pack.BuildUserRegResp(resp, &user)
	}
	return nil
}

func (*CoreService) UserLogin(_ context.Context, req *service.DouyinUserLoginRequest, resp *service.DouyinUserLoginResponse) error {
	var user model.User
	user.Name = req.GetUsername()
	if exist := db.CheckUserExist(user.Name); !exist {
		return errno.UserNotExistErr
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

func (*CoreService) User(_ context.Context, req *service.DouyinUserRequest, resp *service.DouyinUserResponse) error {
	// 查看自己信息时checkUserId == userId
	// 查看他人信息时userId为抖音使用者id checkUserId为发视频作者id
	checkUserId := req.GetUserId()
	userId := utils.GetUserId(req.GetToken())

	// 获取用户信息
	if exist := db.CheckUserIdExist(checkUserId); exist {
		checkUserInfo := db.GetUserInfoById(checkUserId)
		followInfo := db.GetUserFollowInfo(checkUserId, userId)
		pack.BuildUserResp(resp, &checkUserInfo, followInfo)
	} else {
		return errno.UserNotExistErr
	}

	return nil
}

func (*CoreService) PublishAction(_ context.Context, req *service.DouyinPublishActionRequest, resp *service.DouyinPublishActionResponse) error {
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

func (*CoreService) PublishList(_ context.Context, req *service.DouyinPublishListRequest, resp *service.DouyinPublishListResponse) error {
	var (
		videosInfo []*service.Video
		author     service.User
		userId     int64
	)
	// 作者id
	checkId := req.GetUserId()
	// 用户id
	userId = utils.GetUserId(req.GetToken())
	if exist := db.CheckUserIdExist(checkId); !exist {
		return errno.UserNotExistErr
	}
	// 作者信息
	infoById := db.GetUserInfoById(checkId)
	author.Name = &infoById.Name
	url := consts.BackgroundImgUrl
	author.BackgroundImage = &url
	author.Id = &checkId
	author.Avatar = &infoById.Avatar
	// 查询自己时userId == checkId
	followInfo := db.GetUserFollowInfo(checkId, userId)
	author.IsFollow = &followInfo.IsFollow
	author.FollowCount = &followInfo.FollowerCount
	author.FollowerCount = &followInfo.FollowerCount
	// 获取指定作者id 查询自己时userId == checkId
	videos := db.GetVideosByUserId(checkId)
	for i := 0; i < len(videos); i++ {
		// 视频信息
		var videoInfo service.Video
		videoInfo.Id = &videos[i].VideoId
		videoInfo.Title = &videos[i].Title
		videoInfo.PlayUrl = &videos[i].PlayUrl
		videoInfo.CoverUrl = &videos[i].CoverUrl
		authId := videos[i].Author
		// 点赞评论信息
		actionInfo := db.GetActionCount(videos[i].VideoId)
		videoInfo.FavoriteCount = &actionInfo.FavoriteCount
		videoInfo.CommentCount = &actionInfo.CommentCount
		checkFavorite := db.CheckFavorite(userId, authId)
		videoInfo.IsFavorite = &checkFavorite
		// 作者信息
		videoInfo.Author = &author
		// 合并到全部所需信息
		videosInfo = append(videosInfo, &videoInfo)
	}
	pack.BuildPublishListResp(resp, videosInfo)
	return nil
}
