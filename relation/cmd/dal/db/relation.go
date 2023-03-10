package db

import (
	"gorm.io/gorm"
	"relation/cmd/model"
	"relation/pkg/errno"
)

// GetFollowerList
// @Description: 根据作者Id查询粉丝列表
// @auth sinre 2023-02-09 22:26:46
// @param checkId 作者id
// @return followers 粉丝列表
func GetFollowerList(checkId int64) (followers []model.Follow) {
	if err := DB.Where("follow_id = ?", checkId).Find(&followers).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return followers
}

// GetUserInfoById
// @Description: 根据id获取用户信息
// @auth sinre 2023-02-09 22:27:25
// @param id 用户id
// @return user 用户信息
func GetUserInfoById(id int64) (user model.User) {
	if err := DB.Omit("password").Find(&user, "user_id = ?", id).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return user
}

// GetUserFollowInfo
// @Description: 获取用户关注数、粉丝数及关注情况
// @auth sinre 2023-02-09 22:28:41
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

// GetFollowerFriendList
// @Description: 从关注列表内查询好友列表
// @auth sinre 2023-02-09 23:37:31
func GetFollowerFriendList(userId int64) []model.Follow {
	var follows []model.Follow
	var followIds []int64
	// 查询用户的关注列表
	if err := DB.Where("follower_id = ?", userId).Find(&follows).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	// 获取用户关注的id
	for i := 0; i < len(follows); i++ {
		followIds = append(followIds, follows[i].FollowId)
	}
	// 查询用户关注的id 是否已关注该用户 得到朋友
	DB.Where("follow_id = ? AND follower_id IN ?", userId, followIds).Find(&follows)
	return follows
}

// GetLastMessage
// @Description: 获取聊天的最后一条消息
// @auth sinre 2023-02-10 00:53:25
// @param fid 朋友id
// @param uid 用户id
// @return model.Message 消息
func GetLastMessage(fid, uid int64) (mes model.Message) {
	DB.Where("from_user_id = ? OR from_user_id = ?", fid, uid).
		Where("to_user_id = ? OR to_user_id = ?", fid, uid).Order("id desc").Limit(1).Find(&mes)
	return mes
}

// CheckFollowExist 检查两人关注信息是否存在
func CheckFollowExist(userId, toUserId int64) bool {
	var fol model.Follow
	if err := DB.Where("follower_id = ? AND follow_id = ?", userId, toUserId).First(&fol).Error; err == gorm.ErrRecordNotFound {
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
	if err := DB.Unscoped().Model(&follow).Where("follower_id = ? AND follow_id = ?", userId, toUserId).Find(&follow).Error; err != nil {
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

// CheckUserExist
// @Description: 检查用户名是否存在
// @param userId 用户Id
// @return bool false为存在
func CheckUserExist(userId int64) bool {
	var user model.User
	var count int64
	if err := DB.Model(&user).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	if count > 0 {
		return false
	} else {
		return true
	}
}

// SendMessage
// @Description: 发送消息
// @param from_user_id 消息发送方id
// @param to_user_id 消息接收方id
// @param content 消息内容
func SendMessage(fromId, toId int64, content string) {
	message := model.Message{
		FromUserId:   fromId,
		ToUserId:     toId,
		Content:      content,
		SenderRead:   1,
		ReceiverRead: 0,
	}
	if err := DB.Create(&message).Error; err != nil {
		panic(errno.DbInsertErr)
	}
}

func GetChats(tuid, usid int64) (msgs model.Message) {
	if err := DB.Where("to_user_id = ? AND from_user_id = ? AND receiver_read = 0", usid, tuid).First(&msgs).Error; err != gorm.ErrRecordNotFound {
		if err := DB.Model(&model.Message{}).Where("id = ?", msgs.Id).Update("receiver_read", "1").Error; err != nil {
			panic(errno.DbUpdateErr)
		}
	}
	return msgs
}

func GetVideosByUserId(checkId int64) (videos []model.Video) {
	if err := DB.Where("author = ?", checkId).Find(&videos).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return videos
}

func GetStarUserById(vid int64) (favorites []model.Favorite) {
	if err := DB.Where("video_id = ?", vid).Find(&favorites).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return favorites
}
func GetFavoriteListByUserId(checkId int64) (favorites []model.Favorite) {
	if err := DB.Where("user_id = ?", checkId).Find(&favorites).Error; err != nil {
		panic(errno.DbSelectErr)
	}
	return favorites
}
