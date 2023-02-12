package core

import (
	"context"
	"relation/cmd/dal/db"
	"relation/cmd/model"
	"relation/cmd/pack"
	"relation/cmd/service"
	"relation/pkg/errno"
	"relation/pkg/utils"
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
		if !db.CheckFollowExist(userId, toUserId, true) {
			// 直接新增关注信息
			db.FollowAction(userId, toUserId)
		} else {
			return errno.FollowErr
		}
		pack.BuildFollowResp(resp)
	} else if actionType == 2 {
		// 存在关注信息
		if Exist := db.CheckFollowExist(userId, toUserId, false); Exist {
			db.DelFollowAction(userId, toUserId)
			pack.BuildFollowResp(resp)
		} else {
			return errno.DelFollowErr
		}
	}
	return nil
}

func (*RelationService) FollowList(_ context.Context, req *service.DouyinRelationFollowListRequest, resp *service.DouyinRelationFollowListResponse) error {
	var userList []*service.User
	// 获取当前登录用户的 id
	userId := utils.GetUserId(req.GetToken())
	checkId := req.GetUserId()
	if !db.CheckUserIdExist(checkId) {
		return errno.UserNotExistErr
	}
	// 返回用户的关注列表（id)
	followIdList := db.GetFollowUserIdList(checkId)
	for _, fid := range followIdList {
		var su service.User
		ub := db.GetUserInfoById(fid)
		uf := db.GetUserFollowInfo(checkId, userId)
		su.Id = &ub.UserId
		su.Name = &ub.Name
		su.FollowCount = &uf.FollowCount
		su.FollowerCount = &uf.FollowerCount
		su.IsFollow = &uf.IsFollow
		userList = append(userList, &su)
	}
	pack.BuildFollowListResp(resp, userList)
	return nil
}

func (*RelationService) FollowerList(_ context.Context, req *service.DouyinRelationFollowerListRequest, resp *service.DouyinRelationFollowerListResponse) error {
	var authorInfo model.Author
	var authorInfos []model.Author
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
	for _, follower := range followers {
		userinfo := db.GetUserInfoById(follower.FollowerId)
		followerInfo := db.GetUserFollowInfo(follower.FollowerId, userId)
		authorInfo.Id = userinfo.UserId
		authorInfo.Name = userinfo.Name
		authorInfo.FollowCount = followerInfo.FollowCount
		authorInfo.FollowerCount = followerInfo.FollowerCount
		authorInfo.IsFollow = followerInfo.IsFollow
		authorInfos = append(authorInfos, authorInfo)
	}
	pack.BuildRelationFollowerListResp(resp, authorInfos)
	return nil
}

func (*RelationService) FriendList(_ context.Context, req *service.DouyinRelationFriendListRequest, resp *service.DouyinRelationFriendListResponse) error {
	var friendInfo model.Friend
	var friendInfos []model.Friend
	userId := req.GetUserId()
	if !db.CheckUserIdExist(userId) {
		return errno.UserNotExistErr
	}
	// 朋友为筛选后的用户的粉丝
	friendList := db.GetFollowerFriendList(userId)
	for _, friend := range friendList {
		userinfo := db.GetUserInfoById(friend.FollowerId)
		followerInfo := db.GetUserFollowInfo(friend.FollowerId, userId)
		message := db.GetLastMessage(friend.FollowerId, userId)
		messageType := message.CheckMessageType(userId)
		friendInfo.Id = userinfo.UserId
		friendInfo.Name = userinfo.Name
		friendInfo.Message = message.Content
		friendInfo.MessageType = messageType
		friendInfo.FollowCount = followerInfo.FollowCount
		friendInfo.FollowerCount = followerInfo.FollowerCount
		friendInfo.IsFollow = followerInfo.IsFollow
		friendInfos = append(friendInfos, friendInfo)
	}
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
	return nil
}
