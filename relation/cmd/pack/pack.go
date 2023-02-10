package pack

import (
	"relation/cmd/model"
	"relation/cmd/service"
	"relation/pkg/consts"
)

func BuildFollowResp(resp *service.DouyinRelationActionResponse) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildFollowListResp(resp *service.DouyinRelationFollowListResponse, userList []*service.User) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.UserList = userList
}

func BuildRelationFollowerListResp(resp *service.DouyinRelationFollowerListResponse, infos []model.Author) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	var users []*service.User
	for i := 0; i < len(infos); i++ {
		info := infos[i]
		user := &service.User{
			Id:            &info.Id,
			Name:          &info.Name,
			FollowCount:   &info.FollowCount,
			FollowerCount: &info.FollowerCount,
			IsFollow:      &info.IsFollow,
		}
		users = append(users, user)
	}
	resp.UserList = users
}

func BuildFriendListResp(resp *service.DouyinRelationFriendListResponse, infos []model.Friend) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	var users []*service.FriendUser
	for i := 0; i < len(infos); i++ {
		info := infos[i]
		user := &service.FriendUser{
			Id:            &info.Id,
			Name:          &info.Name,
			FollowCount:   &info.FollowCount,
			FollowerCount: &info.FollowerCount,
			IsFollow:      &info.IsFollow,
			Message:       &info.Message,
			MsgType:       &info.MessageType,
		}
		users = append(users, user)
	}
	resp.UserList = users
}
