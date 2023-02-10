package core

import (
	"context"
	"relation/cmd/dal/db"
	"relation/cmd/pack"
	"relation/cmd/service"
	"relation/pkg/errno"
	"relation/pkg/utils"
)

func (*RelationService) RelationAction(ctx context.Context, req *service.DouyinRelationActionRequest, out *service.DouyinRelationActionResponse) error {
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
		pack.BuildFollowResp(out)
	} else if actionType == 2 {
		// 存在关注信息
		if exit := db.CheckFollowExit(userId, toUserId, false); exit {
			db.DelFollowAction(userId, toUserId)
			pack.BuildFollowResp(out)
		} else {
			panic(errno.ActionErr)
		}
	}
	return nil
}

func (*RelationService) FollowList(ctx context.Context, req *service.DouyinRelationFollowListRequest, out *service.DouyinRelationFollowListResponse) error {

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
	pack.BuildFollowListResp(out, userList)
	return nil
}

func (*RelationService) FollowerList(ctx context.Context, req *service.DouyinRelationFollowerListRequest, out *service.DouyinRelationFollowerListResponse) error {
	return nil
}
func (*RelationService) FriendList(ctx context.Context, req *service.DouyinRelationFriendListRequest, out *service.DouyinRelationFriendListResponse) error {
	return nil
}
func (*RelationService) MessageAction(ctx context.Context, req *service.DouyinMessageActionRequest, out *service.DouyinMessageActionResponse) error {
	return nil
}
func (*RelationService) MessageChat(ctx context.Context, req *service.DouyinMessageChatRequest, out *service.DouyinMessageChatResponse) error {
	return nil
}
