package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Comment struct {
	Id          int64                 `gorm:"column:id" json:"id"`
	UserId      int64                 `gorm:"column:user_id" json:"user_id"`
	VideoId     int64                 `gorm:"column:video_id" json:"video_id"`
	CommentText string                `gorm:"column:comment_text" json:"comment_text"`
	CommentTime int64                 `gorm:"column:comment_time" json:"comment_time"`
	IsDeleted   soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted" json:"is_deleted"`
}

type Favorite struct {
	Id        int64                 `gorm:"column:id" json:"id"`
	UserId    int64                 `gorm:"column:user_id" json:"user_id"`
	VideoId   int64                 `gorm:"column:video_id" json:"video_id"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted" json:"is_deleted"`
}

type Follow struct {
	Id         int64                 `gorm:"column:id" json:"id"`
	FollowId   int64                 `gorm:"column:follow_id" json:"follow_id"`
	FollowerId int64                 `gorm:"column:follower_id" json:"follower_id"`
	IsDeleted  soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted" json:"is_deleted"`
}

type Message struct {
	Id         int64  `gorm:"column:id" json:"id"`
	FromUserId int    `gorm:"column:from_user_id" json:"from_user_id"`
	ToUserId   int    `gorm:"column:to_user_id" json:"to_user_id"`
	Content    string `gorm:"column:content" json:"content"`
	SendTime   int64  `gorm:"column:send_time" json:"send_time"`
}

type User struct {
	UserId   int64  `gorm:"column:user_id;autoIncrement" json:"user_id"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
}

type Video struct {
	VideoId    int64     `gorm:"column:video_id" json:"video_id"`
	Title      string    `gorm:"column:title" json:"title"`
	Author     int64     `gorm:"column:author" json:"author"`
	PlayUrl    string    `gorm:"column:play_url" json:"play_url"`
	CoverUrl   string    `gorm:"column:cover_url" json:"cover_url"`
	UploadTime time.Time `gorm:"column:upload_time" json:"upload_time"`
}
