// Package model
// @Description: 数据传输对象(DTO)
package model

import (
	"core/pkg/consts"
	"golang.org/x/crypto/bcrypt"
)

// FollowInfo
// @Description: 关注信息
type FollowInfo struct {
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

// ActionInfo
// @Description: 点赞、评论数据
type ActionInfo struct {
	FavoriteCount int64
	CommentCount  int64
}

// Author
// @Description: 作者信息
type Author struct {
	Id              int64
	Name            string
	Avatar          string
	BackgroundImage string
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

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
