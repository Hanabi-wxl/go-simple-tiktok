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

func BuildCommentActionResp(resp *service.DouyinCommentActionResponse, comment *model.Comment, userInfo *model.User, followInfo model.FollowInfo) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	url := consts.BackgroundImgUrl
	if comment != nil && userInfo != nil {
		formatTime := comment.CommentTime.Format("01-02")
		resp.Comment = &service.Comment{
			Id:         &comment.Id,
			Content:    &comment.CommentText,
			CreateDate: &formatTime,
			User: &service.User{
				Id:              &userInfo.UserId,
				Name:            &userInfo.Name,
				Avatar:          &userInfo.Avatar,
				BackgroundImage: &url,
				FollowCount:     &followInfo.FollowCount,
				FollowerCount:   &followInfo.FollowerCount,
				IsFollow:        &followInfo.IsFollow,
			},
		}
	}
}

func BuildCommentListResp(resp *service.DouyinCommentListResponse, comments []*service.Comment) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.CommentList = comments
}
