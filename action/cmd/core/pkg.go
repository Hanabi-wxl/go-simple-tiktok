package core

import (
	"action/cmd/dal/db"
	"action/cmd/dal/redis"
	"action/cmd/pack"
	"action/cmd/service"
	"strconv"
	"sync"
)

type ActionService struct {
}

func getUserCountInfo(userId int64) (totalFav, workCount, starCount int64) {
	suid := strconv.Itoa(int(userId))
	var userWorkIds []int64

	// 判断works中是否存在checkUserId(userId)
	if exist := redis.CheckUserIdExistInWorks(suid); exist {
		userWorkIds = redis.GetVideoIdsInWorks(suid)
	} else {
		// 不存在则添加key
		redis.CreateUserIdInWorks(suid)
		redis.AddExpireInWorks(suid)
		videos := db.GetVideosByUserId(userId)
		for i := 0; i < len(videos); i++ {
			if ok := redis.AddVideoIdInWorks(suid, videos[i].VideoId); !ok {
				redis.DeleteUserIdInWorks(suid)
			} else {
				userWorkIds = append(userWorkIds, videos[i].VideoId)
			}
		}
	}
	workCount = int64(len(userWorkIds))

	// 判断stars中是否存在videoId
	for i := 0; i < len(userWorkIds); i++ {
		svid := strconv.Itoa(int(userWorkIds[i]))
		if existInStars := redis.CheckVideoIdExistInStars(svid); existInStars {
			// 获取点赞人数
			stars := redis.GetUserIdsInStars(svid)
			totalFav += int64(len(stars))
		} else {
			redis.CreateVideoIdInStars(svid)
			redis.AddExpireInStars(svid)
			// 获取点赞人数
			favorites := db.GetStarUserById(userWorkIds[i])
			for _, user := range favorites {
				if save := redis.AddUserIdInStars(svid, user.UserId); !save {
					// 出现异常删除Key
					redis.DeleteVideoIdInStars(svid)
				}
			}
			totalFav += int64(len(favorites))
		}
	}

	// 判断star中是否存在userId
	if existInStar := redis.CheckUserIdExistInStar(suid); existInStar {
		starVideos := redis.GetVideoIdsInStar(suid)
		starCount = int64(len(starVideos))
	} else {
		// 不存在Key则新增Key
		redis.CreateUserIdInStar(suid)
		// 添加过期时间
		redis.AddExpireInStar(suid)
		// 获取用户所有点赞信息 保存至redis
		favoriteList := db.GetFavoriteListByUserId(userId)
		for _, likeVideoId := range favoriteList {
			if save := redis.AddVideoIdInStar(suid, likeVideoId.VideoId); !save {
				// 如果redis保存出现问题则删掉该Key
				redis.DeleteUserIdInStar(suid)
			}
		}
		starCount = int64(len(favoriteList))
	}
	return totalFav, workCount, starCount
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
	comment := db.GetCommentById(cid)
	// 用户信息
	userInfo := db.GetUserInfoById(comment.UserId)
	// 点赞评论信息
	followInfo := db.GetUserFollowInfo(comment.UserId, userId)
	commentInfo.Id = &comment.Id
	commentInfo.Content = &comment.CommentText
	format := comment.CommentTime.Format("01-02")
	commentInfo.CreateDate = &format

	totalFav, workCount, starCount := getUserCountInfo(userId)
	author := pack.BuildAuthor(userInfo, followInfo, userId, totalFav, workCount, starCount)
	commentInfo.User = author
	(*commentInfos)[index] = &commentInfo
}
