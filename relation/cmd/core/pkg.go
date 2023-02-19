package core

import (
	"relation/cmd/dal/db"
	"relation/cmd/model"
	"relation/cmd/service"
	"relation/pkg/consts"
	"sync"
)

type RelationService struct {
}

func addToFollowList(index int, fid, uid int64, userList *[]*service.User, wg *sync.WaitGroup) {
	defer wg.Done()
	var su service.User
	ub := db.GetUserInfoById(fid)
	uf := db.GetUserFollowInfo(fid, uid)
	su.Id = &ub.UserId
	su.Name = &ub.Name
	su.Avatar = &ub.Avatar
	url := consts.BackgroundImgUrl
	su.BackgroundImage = &url
	su.FollowCount = &uf.FollowCount
	su.FollowerCount = &uf.FollowerCount
	su.IsFollow = &uf.IsFollow
	(*userList)[index] = &su
}
func addToFollowerList(index int, uid int64, follower *model.Follow, authorInfos *[]*service.User, wg *sync.WaitGroup) {
	defer wg.Done()
	var authorInfo service.User
	userinfo := db.GetUserInfoById(follower.FollowerId)
	followerInfo := db.GetUserFollowInfo(follower.FollowerId, uid)
	authorInfo.Id = &userinfo.UserId
	authorInfo.Name = &userinfo.Name
	authorInfo.Avatar = &userinfo.Avatar
	url := consts.BackgroundImgUrl
	authorInfo.BackgroundImage = &url
	authorInfo.FollowCount = &followerInfo.FollowCount
	authorInfo.FollowerCount = &followerInfo.FollowerCount
	authorInfo.IsFollow = &followerInfo.IsFollow
	(*authorInfos)[index] = &authorInfo
}

func addToFriendList(index int, friend model.Follow, uid int64, friendInfos *[]*service.FriendUser, wg *sync.WaitGroup) {
	defer wg.Done()
	var friendInfo service.FriendUser
	userinfo := db.GetUserInfoById(friend.FollowerId)
	followerInfo := db.GetUserFollowInfo(friend.FollowerId, uid)
	message := db.GetLastMessage(friend.FollowerId, uid)
	messageType := message.CheckMessageType(uid)
	friendInfo.Id = &userinfo.UserId
	friendInfo.Name = &userinfo.Name
	friendInfo.Message = &message.Content
	friendInfo.Avatar = &userinfo.Avatar
	friendInfo.MsgType = &messageType
	friendInfo.FollowCount = &followerInfo.FollowCount
	friendInfo.FollowerCount = &followerInfo.FollowerCount
	friendInfo.IsFollow = &followerInfo.IsFollow
	(*friendInfos)[index] = &friendInfo
}
