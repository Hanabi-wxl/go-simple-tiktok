package pack

import (
	"core/cmd/model"
	"core/cmd/service"
	"core/pkg/consts"
)

func BuildUserRegResp(resp *service.DouyinUserRegisterResponse, user *model.User) {
	var token = "token"
	id := &(user.UserId)
	resp.UserId = id

	resp.Token = &token
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildUserLoginResp(resp *service.DouyinUserLoginResponse, user *model.User) {
	var token = "token"

	id := &(user.UserId)
	resp.UserId = id

	resp.Token = &token
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildPublishActionResp(resp *service.DouyinPublishActionResponse) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildUserResp(resp *service.DouyinUserResponse, checkUserInfo *model.User, followInfo model.FollowInfo) {
	var (
		uuser = service.User{
			Id:            &checkUserInfo.UserId,
			Name:          &checkUserInfo.Name,
			FollowCount:   &followInfo.FollowCount,
			FollowerCount: &followInfo.FollowerCount,
			IsFollow:      &followInfo.IsFollow,
		}
	)
	resp.User = &uuser
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildFeedResp(resp *service.DouyinFeedResponse, infos []model.VideoInfo, lastTime int64) {
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
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.NextTime = &lastTime
	resp.VideoList = videoInfos
}

func BuildPublishListResp(resp *service.DouyinPublishListResponse, infos []model.VideoInfo) {
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
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.VideoList = videoInfos
}
