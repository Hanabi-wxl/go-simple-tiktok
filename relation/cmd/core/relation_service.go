package core

import (
	"context"
	"relation/cmd/dal/db"
	"relation/cmd/pack"
	"relation/cmd/service"
	"relation/pkg/errno"
	"relation/pkg/utils"
	"sync"
)

func (*RelationService) RelationAction(_ context.Context, req *service.DouyinRelationActionRequest, resp *service.DouyinRelationActionResponse) error {
	actionType := req.GetActionType()
	toUserId := req.GetToUserId()
	userId := utils.GetUserId(req.GetToken())
	if actionType == 1 {
		if !db.CheckUserIdExist(toUserId) {
			return errno.UserNotExistErr
		}
		// 不存在关注信息
		if !db.CheckFollowExist(userId, toUserId) {
			// 直接新增关注信息
			db.FollowAction(userId, toUserId)
		} else {
			return errno.FollowErr
		}
		pack.BuildFollowResp(resp)
	} else if actionType == 2 {
		// 存在关注信息
		if exist := db.CheckFollowExist(userId, toUserId); exist {
			db.DelFollowAction(userId, toUserId)
			pack.BuildFollowResp(resp)
		} else {
			return errno.DelFollowErr
		}
	}
	return nil
}

func (*RelationService) FollowList(_ context.Context, req *service.DouyinRelationFollowListRequest, resp *service.DouyinRelationFollowListResponse) error {
	// 获取当前登录用户的 id
	userId := utils.GetUserId(req.GetToken())
	checkId := req.GetUserId()
	if !db.CheckUserIdExist(checkId) {
		return errno.UserNotExistErr
	}
	// 返回用户的关注列表（id)
	followIdList := db.GetFollowUserIdList(checkId)
	followLen := len(followIdList)
	var wg sync.WaitGroup
	wg.Add(followLen)
	userList := make([]*service.User, followLen)
	for index, fid := range followIdList {
		go addToFollowList(index, fid, userId, &userList, &wg)
	}
	wg.Wait()
	pack.BuildFollowListResp(resp, userList)
	return nil
}

func (*RelationService) FollowerList(_ context.Context, req *service.DouyinRelationFollowerListRequest, resp *service.DouyinRelationFollowerListResponse) error {
	var userId int64
	claims, _ := utils.ParseToken(req.GetToken())
	// 未登录可查看他人粉丝 默认用户Id为0
	if claims != nil {
		userId = claims.UserId
	}
	checkId := req.GetUserId()
	if !db.CheckUserIdExist(checkId) {
		return errno.UserNotExistErr
	}
	// 获取粉丝
	followers := db.GetFollowerList(checkId)
	ferLen := len(followers)
	var wg sync.WaitGroup
	wg.Add(ferLen)
	authorInfos := make([]*service.User, ferLen)
	for i := 0; i < len(followers); i++ {
		go addToFollowerList(i, userId, &followers[i], &authorInfos, &wg)
	}
	wg.Wait()
	pack.BuildRelationFollowerListResp(resp, authorInfos)
	return nil
}

func (*RelationService) FriendList(_ context.Context, req *service.DouyinRelationFriendListRequest, resp *service.DouyinRelationFriendListResponse) error {
	userId := req.GetUserId()
	if !db.CheckUserIdExist(userId) {
		return errno.UserNotExistErr
	}
	// 朋友为筛选后的用户的粉丝
	friendList := db.GetFollowerFriendList(userId)
	friendLen := len(friendList)
	var wg sync.WaitGroup
	wg.Add(friendLen)
	friendInfos := make([]*service.FriendUser, friendLen)
	for i := 0; i < len(friendList); i++ {
		go addToFriendList(i, friendList[i], userId, &friendInfos, &wg)
	}
	wg.Wait()
	pack.BuildFriendListResp(resp, friendInfos)
	return nil
}

func (*RelationService) MessageAction(_ context.Context, req *service.DouyinMessageActionRequest, resp *service.DouyinMessageActionResponse) error {
	actionType := req.GetActionType()
	toUserId := req.GetToUserId()
	userId := utils.GetUserId(req.GetToken())
	content := req.GetContent()
	// 发送消息操作
	if actionType == 1 {
		// 检查发送对象是否存在，若 exit = true，则对象不存在，返回报错；若为 false，则发送消息
		if exit := db.CheckUserExist(toUserId); exit {
			return errno.UserNotExistErr
		} else {
			db.SendMessage(userId, toUserId, content)
			pack.BuildMessageActionResp(resp)
		}
	}

	return nil
}

func (*RelationService) MessageChat(_ context.Context, req *service.DouyinMessageChatRequest, resp *service.DouyinMessageChatResponse) error {
	var chatss []*service.Message
	toUserId := req.GetToUserId()
	userId := utils.GetUserId(req.GetToken())
	if !db.CheckUserIdExist(toUserId) {
		return errno.UserNotExistErr
	}
	hisChats := db.GetChats(toUserId, userId)
	var chat service.Message
	if hisChats.Id != 0 {
		chat.Id = &hisChats.Id
		chat.ToUserId = &hisChats.ToUserId
		chat.FromUserId = &hisChats.FromUserId
		chat.Content = &hisChats.Content
		milli := hisChats.SendTime.UnixMilli()
		chat.CreateTime = &milli
		chatss = append(chatss, &chat)
	}
	pack.BuildMessageChatResp(resp, chatss)
	return nil
}
