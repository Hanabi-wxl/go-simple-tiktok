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

func BuildUserResp(resp *service.DouyinUserResponse, user *service.User) {
	resp.User = user
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

func BuildAuthor(infoById model.User, followInfo model.FollowInfo, checkId, totalFav, workCount, starCount int64) *service.User {
	var author service.User
	author.Signature = &infoById.Signature
	author.TotalFavorited = &totalFav
	author.WorkCount = &workCount
	author.FavoriteCount = &starCount
	author.Name = &infoById.Name
	url := consts.BackgroundImgUrl
	author.BackgroundImage = &url
	author.Id = &checkId
	author.Avatar = &infoById.Avatar
	author.IsFollow = &followInfo.IsFollow
	author.FollowCount = &followInfo.FollowerCount
	author.FollowerCount = &followInfo.FollowerCount
	return &author
}
