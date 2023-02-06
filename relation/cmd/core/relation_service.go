package core

import (
	"context"
	"relation/cmd/service"
)

func (*RelationService) RelationAction(ctx context.Context, req *service.DouyinRelationActionRequest, out *service.DouyinRelationActionResponse) error {
	return nil
}
func (*RelationService) FollowList(ctx context.Context, req *service.DouyinRelationFollowListRequest, out *service.DouyinRelationFollowListResponse) error {
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
