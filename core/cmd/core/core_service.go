package core

import (
	"context"
	"core/cmd/dal/db"
	"core/cmd/dal/redis"
	"core/cmd/model"
	"core/cmd/pack"
	"core/cmd/service"
	"core/pkg/consts"
	"core/pkg/errno"
	"core/pkg/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func (*CoreService) Feed(_ context.Context, req *service.DouyinFeedRequest, resp *service.DouyinFeedResponse) error {
	var (
		videos   []model.Video
		lastTime int64
		userId   int64
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
	videoLen := len(videos)
	var wg sync.WaitGroup
	wg.Add(videoLen)
	videoInfos := make([]*service.Video, videoLen)
	for i := 0; i < len(videos); i++ {
		go addToFeedVideo(i, videos[i], userId, &videoInfos, &wg)
	}
	wg.Wait()
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

	totalFav, workCount, starCount := getUserCountInfo(checkUserId)

	// 获取用户信息
	if exist := db.CheckUserIdExist(checkUserId); exist {
		checkUserInfo := db.GetUserInfoById(checkUserId)
		followInfo := db.GetUserFollowInfo(checkUserId, userId)
		url := consts.BackgroundImgUrl
		user := service.User{
			Id:              &checkUserInfo.UserId,
			Name:            &checkUserInfo.Name,
			Avatar:          &checkUserInfo.Avatar,
			Signature:       &checkUserInfo.Signature,
			WorkCount:       &workCount,
			FavoriteCount:   &starCount,
			TotalFavorited:  &totalFav,
			BackgroundImage: &url,
			FollowCount:     &followInfo.FollowCount,
			FollowerCount:   &followInfo.FollowerCount,
			IsFollow:        &followInfo.IsFollow,
		}
		pack.BuildUserResp(resp, &user)
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
	info := db.CreateFileInfo(video)
	suid := strconv.Itoa(int(info.Author))
	if inWorks := redis.CheckUserIdExistInWorks(suid); inWorks {
		redis.AddVideoIdInWorks(suid, info.VideoId)
	} else {
		redis.CreateUserIdInWorks(suid)
		redis.AddExpireInWorks(suid)
		videos := db.GetVideosByUserId(info.Author)
		for i := 0; i < len(videos); i++ {
			if ok := redis.AddVideoIdInWorks(suid, videos[i].VideoId); !ok {
				redis.DeleteUserIdInWorks(suid)
			}
		}
	}
	pack.BuildPublishActionResp(resp)
	return nil
}

func (*CoreService) PublishList(_ context.Context, req *service.DouyinPublishListRequest, resp *service.DouyinPublishListResponse) error {
	var videoIds []int64
	var userId int64
	// 作者id
	checkId := req.GetUserId()
	suid := strconv.Itoa(int(checkId))
	// 用户id
	userId = utils.GetUserId(req.GetToken())
	if exist := db.CheckUserIdExist(checkId); !exist {
		return errno.UserNotExistErr
	}
	// 作者信息
	infoById := db.GetUserInfoById(checkId)
	totalFav, workCount, starCount := getUserCountInfo(checkId)
	// 查询自己时userId == checkId
	followInfo := db.GetUserFollowInfo(checkId, userId)
	author := pack.BuildAuthor(infoById, followInfo, checkId, totalFav, workCount, starCount)
	videoIds = redis.GetVideoIdsInWorks(suid)
	videoLen := len(videoIds)
	var wg sync.WaitGroup
	wg.Add(videoLen)
	videosInfo := make([]*service.Video, videoLen)
	for i := 0; i < videoLen; i++ {
		go addToPublishList(i, videoIds[i], userId, author, &videosInfo, &wg)
	}
	wg.Wait()
	pack.BuildPublishListResp(resp, videosInfo)
	return nil
}

func (*CoreService) GetUserListInfo(_ context.Context, req *service.UserListInfoReq, resp *service.UserListInfoResp) error {
	userId := req.GetUserId()
	suid := strconv.Itoa(int(userId))
	var userWorkIds []int64
	var workCount, totalFav, starCount int64
	// 判断works中是否存在checkUserId(userId)
	if exist := redis.CheckUserIdExistInWorks(suid); exist {
		userWorkIds = redis.GetVideoIdsInWorks(suid)
	} else {
		// 不存在则添加key
		redis.CreateUserIdInWorks(suid)
		redis.AddExpireInWorks(suid)
		videos := db.GetVideosByUserId(userId)
		for i := 0; i < len(videos); i++ {
			if ok := redis.AddVideoIdInWorks(suid, videos[i].VideoId); !ok {
				redis.DeleteUserIdInWorks(suid)
			} else {
				userWorkIds = append(userWorkIds, videos[i].VideoId)
			}
		}
	}
	workCount = int64(len(userWorkIds))

	// 判断stars中是否存在videoId
	for i := 0; i < len(userWorkIds); i++ {
		svid := strconv.Itoa(int(userWorkIds[i]))
		if existInStars := redis.CheckVideoIdExistInStars(svid); existInStars {
			// 获取点赞人数
			stars := redis.GetUserIdsInStars(svid)
			totalFav += int64(len(stars))
		} else {
			redis.CreateVideoIdInStars(svid)
			redis.AddExpireInStars(svid)
			// 获取点赞人数
			favorites := db.GetStarUserById(userWorkIds[i])
			for _, user := range favorites {
				if save := redis.AddUserIdInStars(svid, user.UserId); !save {
					// 出现异常删除Key
					redis.DeleteVideoIdInStars(svid)
				}
			}
			totalFav += int64(len(favorites))
		}
	}

	// 判断star中是否存在userId
	if existInStar := redis.CheckUserIdExistInStar(suid); existInStar {
		starVideos := redis.GetVideoIdsInStar(suid)
		starCount = int64(len(starVideos))
	} else {
		// 不存在Key则新增Key
		redis.CreateUserIdInStar(suid)
		// 添加过期时间
		redis.AddExpireInStar(suid)
		// 获取用户所有点赞信息 保存至redis
		favoriteList := db.GetFavoriteListByUserId(userId)
		for _, likeVideoId := range favoriteList {
			if save := redis.AddVideoIdInStar(suid, likeVideoId.VideoId); !save {
				// 如果redis保存出现问题则删掉该Key
				redis.DeleteUserIdInStar(suid)
			}
		}
		starCount = int64(len(favoriteList))
	}
	resp.TotalFavorited = &totalFav
	resp.WorkCount = &workCount
	resp.FavoriteCount = &starCount
	return nil
}
