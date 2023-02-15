package pack

import (
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

func BuildRelationFollowerListResp(resp *service.DouyinRelationFollowerListResponse, infos []*service.User) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.UserList = infos
}

func BuildFriendListResp(resp *service.DouyinRelationFriendListResponse, infos []*service.FriendUser) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.UserList = infos
}

func BuildMessageActionResp(resp *service.DouyinMessageActionResponse) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
}

func BuildMessageChatResp(resp *service.DouyinMessageChatResponse, chats []*service.Message) {
	resp.StatusCode = &consts.DefaultCode
	resp.StatusMsg = &consts.DefaultMsg
	resp.MessageList = chats
}
