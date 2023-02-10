// Package model
// @Description: 数据传输对象(DTO)
package model

// FollowInfo
// @Description: 关注信息
type FollowInfo struct {
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}
