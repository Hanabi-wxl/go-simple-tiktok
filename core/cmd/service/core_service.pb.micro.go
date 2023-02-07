// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: core_service.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CoreService service

func NewCoreServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CoreService service

type CoreService interface {
	Feed(ctx context.Context, in *DouyinFeedRequest, opts ...client.CallOption) (*DouyinFeedResponse, error)
	UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...client.CallOption) (*DouyinUserRegisterResponse, error)
	UserLogin(ctx context.Context, in *DouyinUserLoginRequest, opts ...client.CallOption) (*DouyinUserLoginResponse, error)
	User(ctx context.Context, in *DouyinUserRequest, opts ...client.CallOption) (*DouyinUserResponse, error)
	PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...client.CallOption) (*DouyinPublishActionResponse, error)
	PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...client.CallOption) (*DouyinPublishListResponse, error)
}

type coreService struct {
	c    client.Client
	name string
}

func NewCoreService(name string, c client.Client) CoreService {
	return &coreService{
		c:    c,
		name: name,
	}
}

func (c *coreService) Feed(ctx context.Context, in *DouyinFeedRequest, opts ...client.CallOption) (*DouyinFeedResponse, error) {
	req := c.c.NewRequest(c.name, "CoreService.Feed", in)
	out := new(DouyinFeedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreService) UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...client.CallOption) (*DouyinUserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "CoreService.UserRegister", in)
	out := new(DouyinUserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreService) UserLogin(ctx context.Context, in *DouyinUserLoginRequest, opts ...client.CallOption) (*DouyinUserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "CoreService.UserLogin", in)
	out := new(DouyinUserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreService) User(ctx context.Context, in *DouyinUserRequest, opts ...client.CallOption) (*DouyinUserResponse, error) {
	req := c.c.NewRequest(c.name, "CoreService.User", in)
	out := new(DouyinUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreService) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...client.CallOption) (*DouyinPublishActionResponse, error) {
	req := c.c.NewRequest(c.name, "CoreService.PublishAction", in)
	out := new(DouyinPublishActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreService) PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...client.CallOption) (*DouyinPublishListResponse, error) {
	req := c.c.NewRequest(c.name, "CoreService.PublishList", in)
	out := new(DouyinPublishListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CoreService service

type CoreServiceHandler interface {
	Feed(context.Context, *DouyinFeedRequest, *DouyinFeedResponse) error
	UserRegister(context.Context, *DouyinUserRegisterRequest, *DouyinUserRegisterResponse) error
	UserLogin(context.Context, *DouyinUserLoginRequest, *DouyinUserLoginResponse) error
	User(context.Context, *DouyinUserRequest, *DouyinUserResponse) error
	PublishAction(context.Context, *DouyinPublishActionRequest, *DouyinPublishActionResponse) error
	PublishList(context.Context, *DouyinPublishListRequest, *DouyinPublishListResponse) error
}

func RegisterCoreServiceHandler(s server.Server, hdlr CoreServiceHandler, opts ...server.HandlerOption) error {
	type coreService interface {
		Feed(ctx context.Context, in *DouyinFeedRequest, out *DouyinFeedResponse) error
		UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, out *DouyinUserRegisterResponse) error
		UserLogin(ctx context.Context, in *DouyinUserLoginRequest, out *DouyinUserLoginResponse) error
		User(ctx context.Context, in *DouyinUserRequest, out *DouyinUserResponse) error
		PublishAction(ctx context.Context, in *DouyinPublishActionRequest, out *DouyinPublishActionResponse) error
		PublishList(ctx context.Context, in *DouyinPublishListRequest, out *DouyinPublishListResponse) error
	}
	type CoreService struct {
		coreService
	}
	h := &coreServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CoreService{h}, opts...))
}

type coreServiceHandler struct {
	CoreServiceHandler
}

func (h *coreServiceHandler) Feed(ctx context.Context, in *DouyinFeedRequest, out *DouyinFeedResponse) error {
	return h.CoreServiceHandler.Feed(ctx, in, out)
}

func (h *coreServiceHandler) UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, out *DouyinUserRegisterResponse) error {
	return h.CoreServiceHandler.UserRegister(ctx, in, out)
}

func (h *coreServiceHandler) UserLogin(ctx context.Context, in *DouyinUserLoginRequest, out *DouyinUserLoginResponse) error {
	return h.CoreServiceHandler.UserLogin(ctx, in, out)
}

func (h *coreServiceHandler) User(ctx context.Context, in *DouyinUserRequest, out *DouyinUserResponse) error {
	return h.CoreServiceHandler.User(ctx, in, out)
}

func (h *coreServiceHandler) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, out *DouyinPublishActionResponse) error {
	return h.CoreServiceHandler.PublishAction(ctx, in, out)
}

func (h *coreServiceHandler) PublishList(ctx context.Context, in *DouyinPublishListRequest, out *DouyinPublishListResponse) error {
	return h.CoreServiceHandler.PublishList(ctx, in, out)
}