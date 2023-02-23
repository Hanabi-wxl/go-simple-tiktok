package pack

import (
	"action/cmd/model"
	"action/cmd/service"
	"action/pkg/consts"
)

func BuildFavoriteActionResp(resp *service.DouyinFavoriteActionResponse) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildFavoriteListResp(resp *service.DouyinFavoriteListResponse, infos []*service.Video) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.VideoList = infos
}

func BuildCommentActionResp(resp *service.DouyinCommentActionResponse, comment *model.Comment, userInfo *service.User) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	if comment != nil && userInfo != nil {
		formatTime := comment.CommentTime.Format("01-02")
		resp.Comment = &service.Comment{
			Id:         &comment.Id,
			Content:    &comment.CommentText,
			CreateDate: &formatTime,
			User:       userInfo,
		}
	}
}

func BuildCommentListResp(resp *service.DouyinCommentListResponse, comments []*service.Comment) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.CommentList = comments
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
