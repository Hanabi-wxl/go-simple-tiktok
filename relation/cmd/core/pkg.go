package core

type RelationService struct {
}

// FollowInfo
// @Description: 关注信息
type FollowInfo struct {
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}
