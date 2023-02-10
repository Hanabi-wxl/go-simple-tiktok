package db

import (
	"action/cmd/model"
	"action/pkg/errno"
	"time"
)

// CheckVideoExit
// @Description: 根据id检查video是否存在
// @auth sinre 2023-02-09 16:28:33
// @param vid 视频id
// @return bool 返回结果
func CheckVideoExit(vid int64) bool {
	var count int64
	if err := DB.Model(&model.Video{}).Where("video_id", vid).Count(&count).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if count == 1 {
		return true
	}
	return false
}

// CreateFavorite
// @Description: 点赞
// @auth sinre 2023-02-09 16:29:03
// @param vid 视频id
// @param uid 用户id
func CreateFavorite(vid, uid int64) {
	fav := model.Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	if err := DB.Create(&fav).Error; err != nil {
		panic(errno.DbInsertErr)
	}

}

// DeleteFavorite
// @Description: 取消点赞
// @auth sinre 2023-02-09 16:29:46
// @param vid 视频id
// @param uid 用户id
func DeleteFavorite(vid, uid int64) {
	fav := model.Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	if err := DB.Where("user_id = ? AND video_id = ?", uid, vid).Delete(&fav).Error; err != nil {
		panic(errno.DbUpdateErr)
	}
}

// CheckFavorite
// @Description: 检查是否已点赞
// @auth sinre 2023-02-09 16:29:59
// @param userId 用户id
// @param videoId 视频id
// @return bool 返回结果
func CheckFavorite(userId, videoId int64) bool {
	var favorite model.Favorite
	if err := DB.Where("user_id = ? AND video_id = ?", userId, videoId).Find(&favorite).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if favorite.UserId == 0 {
		return false
	}
	return true
}

// GetUserFollowInfo
// @Description: 获取用户关注数、粉丝数及关注情况
// @auth sinre 2023-02-09 16:30:23
// @param checkUserId 视频作者id
// @param userId 用户id
// @return model.FollowInfo 关注粉丝等数据
func GetUserFollowInfo(checkUserId, userId int64) model.FollowInfo {
	var (
		followModel   model.Follow
		followInfo    model.FollowInfo
		followCount   int64
		followerCount int64
		checkFollow   int64
	)
	// 关注数
	if err := DB.Model(&followModel).Where("follower_id = ?", checkUserId).Count(&followCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	// 粉丝数
	if err := DB.Model(&followModel).Where("follow_id = ?", checkUserId).Count(&followerCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	// 是否已关注
	if err := DB.Model(&followModel).Where("follow_id = ? AND follower_id = ?", checkUserId, userId).
		Count(&checkFollow).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if checkFollow == 1 {
		followInfo.IsFollow = true
	}
	followInfo.FollowCount = followCount
	followInfo.FollowerCount = followerCount
	return followInfo
}

// GetUserInfoById
// @Description: 根据id获取用户信息
// @auth sinre 2023-02-09 16:31:43
// @param id 用户id
// @return user 用户信息
func GetUserInfoById(id int64) (user model.User) {
	if err := DB.Omit("password").Find(&user, "user_id = ?", id).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return user
}

// GetActionCount
// @Description: 获取点赞、评论数
// @auth sinre 2023-02-09 16:32:07
// @param id 视频id
// @return model.ActionInfo 点赞评论等数据
func GetActionCount(id int64) model.ActionInfo {
	var (
		favorite      model.Favorite
		favoriteCount int64
		comment       model.Comment
		commentCount  int64
	)
	if err := DB.Model(&favorite).Where("video_id", id).Count(&favoriteCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if err := DB.Model(&comment).Where("video_id", id).Count(&commentCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return model.ActionInfo{
		CommentCount:  commentCount,
		FavoriteCount: favoriteCount,
	}
}

// GetFavoriteListByUserId
// @Description: 根据id获取点赞列表
// @auth sinre 2023-02-09 16:32:48
// @param checkId 用户id
// @return favorites 点赞列表
func GetFavoriteListByUserId(checkId int64) (favorites []model.Favorite) {
	if err := DB.Where("user_id = ?", checkId).Find(&favorites).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return favorites
}

// GetFavoriteVideos
// @Description: 获取所有点赞的视频信息
// @auth sinre 2023-02-09 16:33:27
// @param vids 视频id
// @return videos 视频信息
func GetFavoriteVideos(vids []int64) (videos []model.Video) {
	if err := DB.Where("video_id IN ?", vids).Find(&videos).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return videos
}

// CreateComment
// @Description: 评论
// @auth sinre 2023-02-09 16:33:59
// @param userId 用户id
// @param videoId 视频id
// @param comment 评论内容
// @return model.Comment 评论详情
func CreateComment(userId int64, videoId int64, comment string) model.Comment {
	commentModel := model.Comment{
		UserId:      userId,
		VideoId:     videoId,
		CommentText: comment,
		CommentTime: time.Now(),
	}
	if err := DB.Create(&commentModel).Error; err != nil {
		panic(errno.DbInsertErr)
	}
	return commentModel
}

// DeleteComment
// @Description: 删除评论
// @auth sinre 2023-02-09 16:34:34
// @param userId 用户id
// @param videoId 视频id
// @param commentId 评论id
func DeleteComment(userId int64, videoId int64, commentId int64) {
	commentModel := model.Comment{
		Id:      commentId,
		UserId:  userId,
		VideoId: videoId,
	}
	if err := DB.Delete(&commentModel).Error; err != nil {
		panic(errno.DbUpdateErr)
	}
}

// GetCommentList
// @Description: 获取评论列表
// @auth sinre 2023-02-09 16:34:50
// @param vid 视频id
// @return comments 评论列表
func GetCommentList(vid int64) (comments []model.Comment) {
	if err := DB.Where("video_id = ?", vid).Find(&comments).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return comments
}