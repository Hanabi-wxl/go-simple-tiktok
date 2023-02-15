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
	resp.Avatar = &user.Avatar
	url := consts.BackgroundImgUrl
	resp.BackgroundImage = &url
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildUserLoginResp(resp *service.DouyinUserLoginResponse, user *model.User) {
	var token = "token"

	id := &(user.UserId)
	resp.UserId = id
	url := consts.BackgroundImgUrl
	resp.Token = &token
	resp.Avatar = &user.Avatar
	resp.BackgroundImage = &url
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildPublishActionResp(resp *service.DouyinPublishActionResponse) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildUserResp(resp *service.DouyinUserResponse, checkUserInfo *model.User, followInfo model.FollowInfo) {
	url := consts.BackgroundImgUrl
	uuser := service.User{
		Id:              &checkUserInfo.UserId,
		Name:            &checkUserInfo.Name,
		Avatar:          &checkUserInfo.Avatar,
		BackgroundImage: &url,
		FollowCount:     &followInfo.FollowCount,
		FollowerCount:   &followInfo.FollowerCount,
		IsFollow:        &followInfo.IsFollow,
	}

	resp.User = &uuser
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildFeedResp(resp *service.DouyinFeedResponse, infos []*service.Video, lastTime int64) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.NextTime = &lastTime
	resp.VideoList = infos
}

func BuildPublishListResp(resp *service.DouyinPublishListResponse, infos []*service.Video) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.VideoList = infos
}
