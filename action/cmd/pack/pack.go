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

func BuildFavoriteListResp(resp *service.DouyinFavoriteListResponse, infos []model.VideoInfo) {
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

func BuildCommentActionResp(resp *service.DouyinCommentActionResponse, comment model.Comment, userInfo model.User, followInfo model.FollowInfo) {
	formatTime := comment.CommentTime.Format("01-02")
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.Comment = &service.Comment{
		Id:         &comment.Id,
		Content:    &comment.CommentText,
		CreateDate: &formatTime,
		User: &service.User{
			Id:            &userInfo.UserId,
			Name:          &userInfo.Name,
			FollowCount:   &followInfo.FollowCount,
			FollowerCount: &followInfo.FollowerCount,
			IsFollow:      &followInfo.IsFollow,
		},
	}
}
func BuildCommentListResp(resp *service.DouyinCommentListResponse, comments []model.CommentInfo) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	var commentInfos []*service.Comment

	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		commentInfo := &service.Comment{
			Id:         &comment.Id,
			CreateDate: &comment.CreateDate,
			Content:    &comment.Content,
			User: &service.User{
				Id:            &comment.User.Id,
				Name:          &comment.User.Name,
				FollowCount:   &comment.User.FollowCount,
				FollowerCount: &comment.User.FollowerCount,
				IsFollow:      &comment.User.IsFollow,
			},
		}
		commentInfos = append(commentInfos, commentInfo)
	}
	resp.CommentList = commentInfos
}
