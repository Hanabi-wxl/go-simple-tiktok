package pack

import (
	"core/cmd/model"
	"core/cmd/service"
)

func BuildUserRegResp(resp *service.DouyinUserRegisterResponse, user *model.User) {
	var (
		token       = "token"
		code  int32 = 0
		msg         = "success"
	)
	id := &(user.UserId)
	resp.UserId = id

	resp.Token = &token
	resp.StatusCode = &code
	resp.StatusMsg = &msg
}

func BuildUserLoginResp(resp *service.DouyinUserLoginResponse, user *model.User) {
	var (
		token       = "token"
		code  int32 = 0
		msg         = "success"
	)
	id := &(user.UserId)
	resp.UserId = id

	resp.Token = &token
	resp.StatusCode = &code
	resp.StatusMsg = &msg
}

func BuildPublishActionResp(resp *service.DouyinPublishActionResponse) {
	var (
		code int32 = 0
		msg        = "success"
	)
	resp.StatusCode = &code
	resp.StatusMsg = &msg
}

func BuildUserResp(resp *service.DouyinUserResponse, checkUserInfo *model.User, followInfo model.FollowInfo) {
	var (
		code  int32 = 0
		msg         = "success"
		uuser       = service.User{
			Id:            &checkUserInfo.UserId,
			Name:          &checkUserInfo.Name,
			FollowCount:   &followInfo.FollowCount,
			FollowerCount: &followInfo.FollowerCount,
			IsFollow:      &followInfo.IsFollow,
		}
	)
	resp.User = &uuser
	resp.StatusCode = &code
	resp.StatusMsg = &msg
}

func BuildFeedRes(resp *service.DouyinFeedResponse, infos []model.VideoInfo, lastTime int64) {
	var (
		code int32 = 0
		msg        = "success"
	)
	var videoInfos []*service.Video
	for i := 0; i < len(infos); i++ {
		info := infos[i]
		videoInfo := &service.Video{
			Id: &info.Id,
			Author: &service.User{
				Id:            &info.Author.Id,
				Name:          &info.Author.Name,
				FollowCount:   &info.Author.FollowCount,
				FollowerCount: &info.Author.FollowerCount,
				IsFollow:      &info.Author.IsFollow,
			},
			PlayUrl:       &info.PlayUrl,
			CoverUrl:      &info.CoverUrl,
			FavoriteCount: &info.FavoriteCount,
			CommentCount:  &info.CommentCount,
			IsFavorite:    &info.IsFavorite,
			Title:         &info.Title,
		}
		videoInfos = append(videoInfos, videoInfo)
	}
	resp.StatusCode = &code
	resp.StatusMsg = &msg
	resp.NextTime = &lastTime
	resp.VideoList = videoInfos
}
