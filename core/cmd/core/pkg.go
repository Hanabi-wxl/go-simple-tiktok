package core

import (
	"core/cmd/dal/db"
	"core/cmd/dal/redis"
	"core/cmd/model"
	"core/cmd/service"
	"core/pkg/consts"
	"strconv"
	"sync"
)

type CoreService struct {
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

func getActionInfo(vid int64) (favCount, comCount int64) {
	svid := strconv.Itoa(int(vid))
	if existInStars := redis.CheckVideoIdExistInStars(svid); existInStars {
		stars := redis.GetUserIdsInStars(svid)
		favCount = int64(len(stars))
	} else {
		redis.CreateVideoIdInStars(svid)
		redis.AddExpireInStars(svid)
		// 获取用户所有点赞信息 保存至redis
		favorites := db.GetStarUserById(vid)
		for _, user := range favorites {
			if save := redis.AddUserIdInStars(svid, user.UserId); !save {
				redis.DeleteVideoIdInStars(svid)
			}
		}
		favCount = int64(len(favorites))
	}

	if inComments := redis.CheckVideoIdInComments(svid); inComments {
		comments := redis.GetCommentIdsInComments(svid)
		comCount = int64(len(comments))
	} else {
		redis.AddVideoIdInComments(svid)
		redis.AddExpireInComments(svid)
		commentList := db.GetCommentList(vid)
		for _, comm := range commentList {
			if save := redis.AddCommentIdInComments(svid, comm.Id); !save {
				redis.DeleteVideoIdInComments(svid)
			}
		}
		comCount = int64(len(commentList))
	}

	return favCount, comCount
}

func addToPublishList(index int, vid, userId int64, author *service.User, videosInfo *[]*service.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	// 视频信息
	var videoInfo service.Video
	videoInfom := db.GetVideoInfoById(vid)
	videoInfo.Id = &videoInfom.VideoId
	videoInfo.Title = &videoInfom.Title
	videoInfo.PlayUrl = &videoInfom.PlayUrl
	videoInfo.CoverUrl = &videoInfom.CoverUrl
	// 点赞评论信息
	favCount, comCount := getActionInfo(videoInfom.VideoId)
	videoInfo.FavoriteCount = &favCount
	videoInfo.CommentCount = &comCount
	checkFavorite := db.CheckFavorite(userId, vid)
	videoInfo.IsFavorite = &checkFavorite
	// 作者信息
	videoInfo.Author = author
	// 合并到全部所需信息
	(*videosInfo)[index] = &videoInfo
}

func addToFeedVideo(index int, video model.Video, userId int64, videoInfos *[]*service.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	var author service.User
	// 视频信息
	var videoInfo service.Video
	videoInfo.Id = &video.VideoId
	videoInfo.Title = &video.Title
	videoInfo.PlayUrl = &video.PlayUrl
	videoInfo.CoverUrl = &video.CoverUrl
	authId := video.Author
	// 点赞评论信息
	favCount, comCount := getActionInfo(video.VideoId)
	videoInfo.FavoriteCount = &favCount
	videoInfo.CommentCount = &comCount
	checkFavorite := db.CheckFavorite(userId, video.VideoId)
	videoInfo.IsFavorite = &checkFavorite
	// 作者信息
	infoById := db.GetUserInfoById(authId)
	totalFav, workCount, starCount := getUserCountInfo(authId)
	author.Name = &infoById.Name
	author.Id = &authId
	author.Avatar = &infoById.Avatar
	author.Signature = &infoById.Signature
	author.TotalFavorited = &totalFav
	author.WorkCount = &workCount
	author.FavoriteCount = &starCount
	url := consts.BackgroundImgUrl
	author.BackgroundImage = &url
	followInfo := db.GetUserFollowInfo(authId, userId)
	author.IsFollow = &followInfo.IsFollow
	author.FollowCount = &followInfo.FollowCount
	author.FollowerCount = &followInfo.FollowerCount
	videoInfo.Author = &author
	// 合并到全部所需信息
	(*videoInfos)[index] = &videoInfo
}
