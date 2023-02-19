package core

import (
	"action/cmd/dal/db"
	"action/cmd/dal/redis"
	"action/cmd/service"
	"action/pkg/consts"
	"strconv"
	"sync"
)

type ActionService struct {
}

func addToFavList(index int, favId, userId int64, author *service.User, videoInfos *[]*service.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	video := db.GetVideoInfoById(favId)
	// 视频信息
	var videoInfo service.Video
	videoInfo.Id = &video.VideoId
	videoInfo.Title = &video.Title
	videoInfo.PlayUrl = &video.PlayUrl
	videoInfo.CoverUrl = &video.CoverUrl
	authId := video.Author
	// 点赞评论信息
	svid := strconv.Itoa(int(video.VideoId))
	stars := redis.GetUserIdsInStars(svid)
	comments := redis.GetCommentIdsInComments(svid)
	fc := int64(len(stars))
	cc := int64(len(comments))
	videoInfo.FavoriteCount = &fc
	videoInfo.CommentCount = &cc
	checkFavorite := db.CheckFavorite(userId, authId)
	videoInfo.IsFavorite = &checkFavorite
	// 作者信息
	videoInfo.Author = author
	// 合并到全部所需信息
	(*videoInfos)[index] = &videoInfo
}

func addToCommentList(index int, cid, userId int64, commentInfos *[]*service.Comment, wg *sync.WaitGroup) {
	defer wg.Done()
	var commentInfo service.Comment
	var user service.User
	comment := db.GetCommentById(cid)
	// 用户信息
	userInfo := db.GetUserInfoById(comment.UserId)
	// 点赞评论信息
	followInfo := db.GetUserFollowInfo(comment.UserId, userId)
	commentInfo.Id = &comment.Id
	commentInfo.Content = &comment.CommentText
	format := comment.CommentTime.Format("01-02")
	commentInfo.CreateDate = &format
	user.Id = &userInfo.UserId
	user.Name = &userInfo.Name
	url := consts.BackgroundImgUrl
	user.Avatar = &userInfo.Avatar
	user.BackgroundImage = &url
	user.FollowCount = &followInfo.FollowCount
	user.FollowerCount = &followInfo.FollowerCount
	user.IsFollow = &followInfo.IsFollow
	commentInfo.User = &user
	(*commentInfos)[index] = &commentInfo
}
