// Package model
// @Description: 数据库实体对象
package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Comment
// @Description: 评论
type Comment struct {
	Id          int64                 `gorm:"column:id;primaryKey" json:"id"`
	UserId      int64                 `gorm:"column:user_id" json:"user_id"`
	VideoId     int64                 `gorm:"column:video_id" json:"video_id"`
	CommentText string                `gorm:"column:comment_text" json:"comment_text"`
	CommentTime time.Time             `gorm:"column:comment_time" json:"comment_time"`
	IsDeleted   soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted" json:"is_deleted"`
}

// Favorite
// @Description: 点赞喜欢
type Favorite struct {
	Id        int64                 `gorm:"column:id;primaryKey" json:"id"`
	UserId    int64                 `gorm:"column:user_id" json:"user_id"`
	VideoId   int64                 `gorm:"column:video_id" json:"video_id"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted" json:"is_deleted"`
}

// Follow
// @Description: 关注
type Follow struct {
	Id         int64                 `gorm:"column:id;primaryKey" json:"id"`
	FollowId   int64                 `gorm:"column:follow_id" json:"follow_id"`
	FollowerId int64                 `gorm:"column:follower_id" json:"follower_id"`
	IsDeleted  soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted" json:"is_deleted"`
}

// User
// @Description: 用户
type User struct {
	UserId    int64  `gorm:"column:user_id;primaryKey" json:"user_id"`
	Password  string `gorm:"column:password" json:"password"`
	Name      string `gorm:"column:name" json:"name"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	Signature string `gorm:"column:signature" json:"signature"`
}

// Video
// @Description: 视频
type Video struct {
	VideoId    int64     `gorm:"column:video_id;primaryKey" json:"video_id"`
	Title      string    `gorm:"column:title" json:"title"`
	Author     int64     `gorm:"column:author" json:"author"`
	PlayUrl    string    `gorm:"column:play_url" json:"play_url"`
	CoverUrl   string    `gorm:"column:cover_url" json:"cover_url"`
	UploadTime time.Time `gorm:"column:upload_time" json:"upload_time"`
}
