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

// ActionInfo
// @Description: 点赞、关注数据
type ActionInfo struct {
	FavoriteCount int64
	CommentCount  int64
}

// Author
// @Description: 作者信息
type Author struct {
	Id   int64
	Name string
	FollowInfo
}

// VideoInfo
// @Description: 视频信息
type VideoInfo struct {
	Id            int64
	Author        Author
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
}

// CommentInfo
// @Description: 评论信息
type CommentInfo struct {
	Id         int64
	User       Author
	Content    string
	CreateDate string
}

type Star struct {
	UserId  int64
	VideoId int64
}
