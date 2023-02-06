package service

import (
	"context"
	"core/cmd/service"
	"fmt"
)

func (*CoreService) Feed(ctx context.Context, req *service.DouyinFeedRequest, resp *service.DouyinFeedResponse) error {
	fmt.Println("qqqq")
	return nil
}
func (*CoreService) UserRegister(ctx context.Context, req *service.DouyinUserRegisterRequest, resp *service.DouyinUserRegisterResponse) error {
	return nil
}
func (*CoreService) UserLogin(ctx context.Context, req *service.DouyinUserLoginRequest, resp *service.DouyinUserLoginResponse) error {
	return nil
}
func (*CoreService) User(ctx context.Context, req *service.DouyinUserRequest, resp *service.DouyinUserResponse) error {
	return nil
}
func (*CoreService) PublishAction(ctx context.Context, req *service.DouyinPublishActionRequest, resp *service.DouyinPublishActionResponse) error {
	return nil
}
func (*CoreService) PublishList(ctx context.Context, req *service.DouyinPublishListRequest, resp *service.DouyinPublishListResponse) error {
	return nil
}
