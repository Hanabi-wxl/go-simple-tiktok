package db

import (
	"core/cmd/model"
	"core/pkg/consts"
	"core/pkg/errno"
	"time"
)

func CheckUserExit(username string) bool {
	var user model.User
	var count int64
	if err := DB.Model(&user).Where("name = ?", username).Count(&count).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

func CreateUser(username, password string) {
	var user model.User
	user.Name = username
	_ = user.SetPassword(password)
	if err := DB.Create(&user).Error; err != nil {
		panic(errno.DbInsertErr)
	}
}

func GetUserInfoByUsername(username string) model.User {
	var user model.User
	if err := DB.Find(&user, "name = ?", username).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return user
}

func CreateFileInfo(video model.Video) {
	video.UploadTime = time.Now()
	if err := DB.Create(&video).Error; err != nil {
		panic(errno.DbInsertErr)
	}
}

func GetUserInfoById(id int64) model.User {
	var user model.User
	if err := DB.Find(&user, "user_id = ?", id).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return user
}

func GetUserFollowInfo(checkUserId, userId int64) model.FollowInfo {
	var (
		followModel   model.Follow
		followInfo    model.FollowInfo
		followCount   int64
		followerCount int64
		checkFollow   int64
	)
	if err := DB.Model(&followModel).Where("follow_id = ?", checkUserId).Count(&followCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if err := DB.Model(&followModel).Where("follower_id = ?", checkUserId).Count(&followerCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
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

func FeedVideos(time time.Time) ([]model.Video, int64) {
	var (
		videoModel model.Video
		videos     []model.Video
		lastTime   int64
	)
	if err := DB.Model(&videoModel).Where("upload_time < ?", time).
		Order("upload_time desc").Limit(consts.VideoLimit).Find(&videos).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if len(videos) > 0 {
		lastTime = videos[consts.VideoLimit-1].UploadTime.UnixMilli()
	}
	return videos, lastTime
}

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
