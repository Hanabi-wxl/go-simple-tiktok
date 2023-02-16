package db

import (
	"github.com/stretchr/testify/assert"
	"relation/cmd/model"
	"testing"
)

func TestCheckFollowExit(t *testing.T) {
	Init()
	isPass := CheckFollowExist(1, 1)
	assert.Equal(t, true, isPass, "测试失败")
}

func TestDelFollowAction(t *testing.T) {
	Init()
	DelFollowAction(10, 20)
	t.Log("DelFollowAction被执行")
}

func TestInit(t *testing.T) {
	Init()
	t.Log("Init()方法被执行,数据库连接成功")
}

func TestFollowActionUpdate(t *testing.T) {
	Init()
	FollowActionUpdate(10, 20)
	t.Log("FollowActionUpdate被执行")
}

func TestGetFollowerList(t *testing.T) {
	Init()
	isPass := GetFollowerList(1)
	a := []model.Follow([]model.Follow{model.Follow{Id: 0, FollowId: 1, FollowerId: 1, IsDeleted: 0x0}})
	assert.Equal(t, a, isPass, "测试获取关注列表失败")
}

func TestGetFollowInfoById(t *testing.T) {
	Init()
	isPass := GetFollowInfoById(1)
	a := model.FollowInfo(model.FollowInfo{FollowCount: 1, FollowerCount: 1, IsFollow: true})
	assert.Equal(t, a, isPass, "测试获取关注列表失败")
}

func TestGetLastMessage(t *testing.T) {
	Init()
	isPass := GetLastMessage(10, 20)
	a := model.Message(model.Message{FromUserId: 1, ToUserId: 1})
	assert.Equal(t, a, isPass, "获取最后一条信息失败")
}

func TestGetUserFollowInfo(t *testing.T) {
	Init()
	isPass := GetUserFollowInfo(10, 20)
	a := model.FollowInfo(model.FollowInfo{FollowCount: 0, FollowerCount: 0, IsFollow: false})
	assert.Equal(t, a, isPass, "测试获取关注以及粉丝数失败")
}

func TestGetUserInfoById(t *testing.T) {
	Init()
	isPass := GetUserInfoById(1)
	a := model.User{UserId: 1, Name: "123"}
	assert.Equal(t, a, isPass, "测试获取用户信息失败")

}

func TestGetFollowUserIdList(t *testing.T) {
	Init()
	isPass := GetFollowUserIdList(1)
	a := []int64([]int64{1})
	assert.Equal(t, a, isPass, "测试获取关注用户失败")
}

func TestGetFollowerFriendList(t *testing.T) {
	Init()
	isPass := GetFollowerFriendList(1)
	a := []model.Follow([]model.Follow{model.Follow{Id: 0, FollowId: 1, FollowerId: 1, IsDeleted: 0x0}})
	assert.Equal(t, a, isPass, "测试获取关注好友列表失败")
}

func TestFollowAction(t *testing.T) {
	Init()
	FollowAction(1, 2)
	t.Log("FollowAction被执行")
}
