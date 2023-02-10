package db

import (
	"relation/cmd/model"
	"relation/pkg/errno"
)

// CheckFollowExit 检查两人关注信息是否存在
func CheckFollowExit(userId, toUserId int64, isUnscoped bool) bool {
	/*
		isUnscoped: 是否添加 Unscoped()
	*/
	var (
		fol   model.Follow
		count int64
	)
	if isUnscoped {
		if err := DB.Unscoped().Model(&fol).Where(&model.Follow{FollowerId: userId, FollowId: toUserId}).Count(&count).Error; err != nil {
			panic(errno.DbSelectErr)
		}
	} else {
		if err := DB.Model(&fol).Where(&model.Follow{FollowerId: userId, FollowId: toUserId}).Count(&count).Error; err != nil {
			panic(errno.DbSelectErr)
		}
	}
	return count > 0
}

// FollowAction 关注某人 (新增)
func FollowAction(userId int64, toUserId int64) {
	var follow model.Follow
	follow.FollowId = toUserId
	follow.FollowerId = userId
	if err := DB.Create(&follow).Error; err != nil {
		panic(errno.FollowErr)
	}
}

// FollowActionUpdate 关注某人 (更改逻辑字段)
func FollowActionUpdate(userId int64, toUserId int64) {
	var follow model.Follow
	if err := DB.Unscoped().Model(&follow).Where(&model.Follow{FollowerId: userId, FollowId: toUserId}).Find(&follow).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	follow.IsDeleted = 0
	if err := DB.Unscoped().Save(&follow).Error; err != nil {
		panic(errno.FollowErr)
	}
}

// DelFollowAction 取消关注某人
func DelFollowAction(userId int64, toUserId int64) {
	var follow model.Follow
	if err := DB.Model(&follow).Where(&model.Follow{FollowerId: userId, FollowId: toUserId}).Delete(&follow).Error; err != nil {
		panic(errno.DelFollowErr)
	}
}

// GetFollowUserIdList 返回用户的关注列表（id)
func GetFollowUserIdList(userId int64) []int64 {
	var (
		followList []model.Follow
		followIds  []int64
	)
	_ = DB.Where(&model.Follow{FollowerId: userId}).Find(&followList)
	for _, fol := range followList {
		followIds = append(followIds, fol.FollowId)
	}
	return followIds
}

// GetUserInfoById 根据id获取用户基本信息

func GetUserInfoById(id int64) (user model.User) {
	if err := DB.Where(&model.User{UserId: id}).Find(&user).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	user.Password = ""
	return user
}

// GetFollowInfoById 根据id获取用户的关注与粉丝信息
func GetFollowInfoById(userId int64) model.FollowInfo {
	var (
		followModel   model.Follow
		followInfo    model.FollowInfo
		followCount   int64
		followerCount int64
	)

	// 关注数
	if err := DB.Model(&followModel).Where("follower_id = ?", userId).Count(&followCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	// 粉丝数
	if err := DB.Model(&followModel).Where("follow_id = ?", userId).Count(&followerCount).Error; err != nil {
		panic(errno.DbSelectErr)
	}

	// 是否已关注 (本就是登录用户的关注列表，直接返回 true)
	followInfo.IsFollow = true
	followInfo.FollowCount = followCount
	followInfo.FollowerCount = followerCount
	return followInfo
}
