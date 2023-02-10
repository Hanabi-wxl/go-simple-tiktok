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
		// 不存在关注信息
		if exit := db.CheckFollowExit(userId, toUserId, true); !exit {
			// 直接新增关注信息
			db.FollowAction(userId, toUserId)
		} else {
			// 通过更改逻辑字段再次关注
			db.FollowActionUpdate(userId, toUserId)
		}
		pack.BuildFollowResp(resp)
	} else if actionType == 2 {
		// 存在关注信息
		if exit := db.CheckFollowExit(userId, toUserId, false); exit {
			db.DelFollowAction(userId, toUserId)
			pack.BuildFollowResp(resp)
		} else {
			panic(errno.ActionErr)
		}
	}
	return nil
}

func (*RelationService) FollowList(_ context.Context, req *service.DouyinRelationFollowListRequest, resp *service.DouyinRelationFollowListResponse) error {
	var userList []*service.User
	// 获取当前登录用户的 id
	userId := utils.GetUserId(req.GetToken())
	// 返回用户的关注列表（id)
	followIdList := db.GetFollowUserIdList(userId)
	for _, fid := range followIdList {
		var su service.User
		ub := db.GetUserInfoById(fid)
		uf := db.GetFollowInfoById(fid)
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
	// "朋友"可理解为筛选后的用户的粉丝
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
	return nil
}
func (*RelationService) MessageChat(_ context.Context, req *service.DouyinMessageChatRequest, resp *service.DouyinMessageChatResponse) error {
	return nil
}
