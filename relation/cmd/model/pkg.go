// Package model
// @Description: 数据传输对象(DTO)
package model

// Author
// @Description: 作者信息
type Author struct {
	Id   int64
	Name string
	FollowInfo
}

// FollowInfo
// @Description: 关注信息
type FollowInfo struct {
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type Friend struct {
	Author
	Message     string
	MessageType int64
}

// CheckMessageType
// @Description: 获取最近的消息类型
// @auth sinre 2023-02-10 01:38:00
// @receiver m 方法对象
// @param uid 用户id
// @return int64 消息类型 0收 1发
func (m *Message) CheckMessageType(uid int64) int64 {
	if m.FromUserId == uid {
		return 1
	}
	return 0
}
