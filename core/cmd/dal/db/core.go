package db

import (
	"core/cmd/model"
	"core/pkg/consts"
	"core/pkg/errno"
	"gorm.io/gorm"
	"time"
)

// CheckUserExist
// @Description: 检查用户名是否存在
// @auth sinre 2023-02-09 16:42:25
// @param username 用户名
// @return bool true为存在
func CheckUserExist(username string) bool {
	var user model.User
	if err := DB.Where("name = ?", username).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// CheckUserIdExist
// @Description: 检查用户id是否存在
// @auth sinre 2023-02-11 18:20:03
// @param id 用户id
// @return bool 存在标志
func CheckUserIdExist(id int64) bool {
	var user model.User
	if err := DB.First(&user, id).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// CreateUser
// @Description: 添加用户 注册
// @auth sinre 2023-02-09 16:42:45
// @param username 用户名
// @param password 密码
func CreateUser(username, password, avatar string) {
	var user model.User
	user.Name = username
	user.Avatar = avatar
	// 生成加密密码
	_ = user.SetPassword(password)
	if err := DB.Create(&user).Error; err != nil {
		panic(errno.DbInsertErr)
	}
}

// GetUserInfoByUsername
// @Description: 根据用户名获取用户信息
// @auth sinre 2023-02-09 16:43:13
// @param username 用户名
// @return user 用户信息
func GetUserInfoByUsername(username string) (user model.User) {
	if err := DB.Find(&user, "name = ?", username).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return user
}

// CreateFileInfo
// @Description: 创建视频文件信息
// @auth sinre 2023-02-09 16:43:40
// @param video 视频信息
func CreateFileInfo(video model.Video) {
	video.UploadTime = time.Now()
	if err := DB.Create(&video).Error; err != nil {
		panic(errno.DbInsertErr)
	}
}

// GetUserInfoById
// @Description: 根据id获取用户信息
// @auth sinre 2023-02-09 16:44:03
// @param id 用户id
// @return user 用户信息
func GetUserInfoById(id int64) (user model.User) {
	if err := DB.Find(&user, "user_id = ?", id).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return user
}

// GetUserFollowInfo
// @Description: 获取用户关注数、粉丝数及关注情况
// @auth sinre 2023-02-09 16:44:58
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
	if checkUserId == userId {
		followInfo.IsFollow = true
	} else {
		if err := DB.Model(&followModel).Where("follow_id = ? AND follower_id = ?", checkUserId, userId).
			Count(&checkFollow).Error; err != nil {
			panic(errno.DbSelectErr)
		}
		if checkFollow == 1 {
			followInfo.IsFollow = true
		}
	}
	followInfo.FollowCount = followCount
	followInfo.FollowerCount = followerCount
	return followInfo
}

// FeedVideos
// @Description: 获取视频流 返回按投稿时间倒序的视频列表
// @auth sinre 2023-02-09 16:45:47
// @param time 时间戳
// @return videos 视频信息
// @return lastTime 较晚发布的视频的时间戳
func FeedVideos(time time.Time) (videos []model.Video, lastTime int64) {
	if err := DB.Where("upload_time < ?", time).Order("upload_time desc").Limit(consts.VideoLimit).Find(&videos).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if len(videos) > 1 {
		lastTime = videos[len(videos)-1].UploadTime.UnixMilli()
	}
	return videos, lastTime
}

// CheckFavorite
// @Description: 检查是否已点赞
// @auth sinre 2023-02-09 16:50:15
// @param userId 用户id
// @param videoId 视频id
// @return bool 点赞信息
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

// GetActionCount
// @Description: 获取互动数据 如点赞数、评论数
// @auth sinre 2023-02-09 16:50:42
// @param id 视频id
// @return model.ActionInfo 互动数据
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

// GetVideosByUserId
// @Description: 获取用户发布视频列表
// @auth sinre 2023-02-09 16:51:11
// @param checkId 用户id
// @return videos 视频信息
func GetVideosByUserId(checkId int64) (videos []model.Video) {
	if err := DB.Where("author = ?", checkId).Find(&videos).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return videos
}
