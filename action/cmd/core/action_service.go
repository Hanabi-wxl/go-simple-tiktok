package core

import (
	"action/cmd/dal/db"
	"action/cmd/dal/redis"
	"action/cmd/model"
	"action/cmd/mq"
	"action/cmd/pack"
	"action/cmd/service"
	"action/pkg/errno"
	"action/pkg/utils"
	"context"
	"encoding/json"
	"strconv"
)

func (*ActionService) FavoriteAction(_ context.Context, req *service.DouyinFavoriteActionRequest, resp *service.DouyinFavoriteActionResponse) error {
	acType := req.GetActionType()
	token := req.GetToken()
	videoId := req.GetVideoId()
	uid := utils.GetUserId(token)
	suid := strconv.Itoa(int(uid))
	svid := strconv.Itoa(int(videoId))
	if exit := db.CheckVideoExit(videoId); exit {
		body, _ := json.Marshal(model.MQStar{UserId: uid, VideoId: videoId})
		if acType == 1 {
			// 维护redis数据库: Star
			// 查询是否存在userId
			if exit1 := redis.CheckUserIdExitInStar(suid); exit1 {
				// 存在Key: userId则创建点赞信息至redis
				redis.AddVideoIdInStar(suid, videoId)
				// 使用mq同步信息至db
				mq.StarAddQue.Publish(body)
			} else {
				// 不存在Key则新增Key
				redis.CreateUserIdInStar(suid)
				// 添加过期时间
				redis.AddExpireInStar(suid)
				// 获取用户所有点赞信息 保存至redis
				favoriteList := db.GetFavoriteListByUserId(uid)
				for _, likeVideoId := range favoriteList {
					if save := redis.AddVideoIdInStar(suid, likeVideoId.VideoId); !save {
						// 如果redis保存出现问题则删掉该Key
						redis.DeleteUserIdInStar(suid)
					}
				}
				// 保存当前点赞信息至redis
				if save := redis.AddVideoIdInStar(suid, videoId); !save {
					// 如果redis保存出现问题则删掉该Key
					redis.DeleteUserIdInStar(suid)
				} else {
					mq.StarAddQue.Publish(body)
				}
			}
			// 维护redis数据库: Stars
			// 查询是否存在videoId
			if exit1 := redis.CheckVideoIdExitInStars(svid); exit1 {
				// 点赞后将userId保存至stars数据库
				redis.AddUserIdInStars(svid, uid)
			} else {
				// 添加Key
				redis.CreateVideoIdInStars(svid)
				// 添加过期时间
				redis.AddExpireInStars(svid)
				// 获取用户所有点赞信息 保存至redis
				favorites := db.GetStarUserById(videoId)
				for _, user := range favorites {
					if save := redis.AddUserIdInStars(svid, user.UserId); !save {
						// 出现异常删除Key
						redis.DeleteVideoIdInStars(svid)
					}
				}
				// 保存当前点赞信息
				if save := redis.AddUserIdInStars(svid, uid); !save {
					redis.DeleteVideoIdInStars(svid)
				}
			}
		} else {
			if exit1 := redis.CheckUserIdExitInStar(suid); exit1 {
				// 删除redis中的点赞信息
				redis.RemoveVideoIdInStar(suid, videoId)
				// 同步删除信息至db
				mq.StarDelQue.Publish(body)
			} else {
				// 不存在Key则新增Key
				redis.CreateUserIdInStar(suid)
				// 添加过期时间
				redis.AddExpireInStar(suid)
				// 获取用户所有点赞信息 保存至redis
				favoriteList := db.GetFavoriteListByUserId(uid)
				for _, likeVideoId := range favoriteList {
					if save := redis.AddVideoIdInStar(suid, likeVideoId.VideoId); !save {
						// 如果redis保存出现问题则删掉该Key
						redis.DeleteUserIdInStar(suid)
					}
				}
				// 删除当前点赞信息
				redis.RemoveVideoIdInStar(suid, videoId)
				// 同步删除信息至db
				mq.StarDelQue.Publish(body)
			}

			if exit1 := redis.CheckVideoIdExitInStars(svid); exit1 {
				redis.RemoveUserIdInStars(svid, uid)
			} else {
				// 添加Key
				redis.CreateVideoIdInStars(suid)
				// 添加过期时间
				redis.AddExpireInStars(suid)
				// 获取用户所有点赞信息 保存至redis
				favorites := db.GetStarUserById(videoId)
				for _, user := range favorites {
					if save := redis.AddUserIdInStars(svid, user.UserId); !save {
						// 出现异常删除Key
						redis.DeleteVideoIdInStars(svid)
					}
				}
				// 删除该点赞信息
				redis.RemoveUserIdInStars(svid, uid)
			}
		}
	} else {
		return errno.VideoNotExit
	}
	pack.BuildFavoriteActionResp(resp)
	return nil
}

func (*ActionService) FavoriteList(_ context.Context, req *service.DouyinFavoriteListRequest, resp *service.DouyinFavoriteListResponse) error {
	var (
		videos     []model.Video
		videoInfos []*service.Video
		videoInfo  service.Video
		author     service.User
		userId     int64
		fids       []int64
	)
	token := req.GetToken()
	checkId := req.GetUserId()
	userId = utils.GetUserId(token)
	suid := strconv.Itoa(int(userId))
	scid := strconv.Itoa(int(checkId))
	// 作者信息
	infoById := db.GetUserInfoById(checkId)
	author.Name = &infoById.Name
	author.Id = &checkId
	// 查询自己时userId == checkId
	followInfo := db.GetUserFollowInfo(checkId, userId)
	author.IsFollow = &followInfo.IsFollow
	author.FollowCount = &followInfo.FollowerCount
	author.FollowerCount = &followInfo.FollowerCount

	if exit := redis.CheckUserIdExitInStar(scid); exit {
		// 查询视频作者喜欢列表
		fids = redis.GetVideoIdsInStar(scid)
		// 查询视频信息
		videos = db.GetFavoriteVideos(fids)
		for _, video := range videos {
			// 视频信息
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
			videoInfo.Author = &author
			// 合并到全部所需信息
			videoInfos = append(videoInfos, &videoInfo)
		}
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
		for i := 0; i < len(favoriteList); i++ {
			fids = append(fids, favoriteList[i].VideoId)
		}
		// 查询视频信息
		videos = db.GetFavoriteVideos(fids)
		for _, video := range videos {
			// 视频信息
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
			videoInfo.Author = &author
			// 合并到全部所需信息
			videoInfos = append(videoInfos, &videoInfo)
		}
	}
	pack.BuildFavoriteListResp(resp, videoInfos)
	return nil
}

func (*ActionService) CommentAction(_ context.Context, req *service.DouyinCommentActionRequest, resp *service.DouyinCommentActionResponse) error {
	var (
		userId     int64
		videoId    = req.GetVideoId()
		svid       = strconv.Itoa(int(videoId))
		acType     = req.GetActionType()
		commentId  = req.GetCommentId()
		comment    model.Comment
		userInfo   model.User
		followInfo model.FollowInfo
	)

	token := req.GetToken()
	userId = utils.GetUserId(token)
	if exit := db.CheckVideoExit(videoId); exit {
		if acType == 1 {
			comment = db.CreateComment(userId, videoId, req.GetCommentText())
			// 判断是否存在key: videoId
			if exit1 := redis.CheckVideoIdInComments(svid); exit1 {
				if save := redis.AddCommentIdInComments(svid, comment.Id); !save {
					redis.DeleteVideoIdInComments(svid)
				}
			} else {
				// 创建key: videoId
				redis.AddVideoIdInComments(svid)
				redis.AddExpireInComments(svid)
				commentList := db.GetCommentList(videoId)
				for i := 0; i < len(commentList); i++ {
					if save := redis.AddCommentIdInComments(svid, commentList[i].Id); !save {
						redis.DeleteVideoIdInComments(svid)
					}
				}
			}
			userInfo = db.GetUserInfoById(userId)
			followInfo = db.GetUserFollowInfo(userId, userId)
			pack.BuildCommentActionResp(resp, comment, userInfo, followInfo)
		} else if acType == 2 {
			// 判断是否存在key: videoId
			body, _ := json.Marshal(&model.MQComment{CommentId: commentId})
			if exit1 := redis.CheckVideoIdInComments(svid); exit1 {
				redis.RemoveCommentIdInComments(svid, commentId)
				mq.CommentDelQue.Publish(body)
			} else {
				// 创建key: videoId
				redis.AddVideoIdInComments(svid)
				redis.AddExpireInComments(svid)
				commentList := db.GetCommentList(videoId)
				for i := 0; i < len(commentList); i++ {
					if save := redis.AddCommentIdInComments(svid, commentList[i].Id); !save {
						redis.DeleteVideoIdInComments(svid)
					}
				}
				mq.CommentDelQue.Publish(body)
			}
			pack.BuildCommentActionResp(resp, comment, userInfo, followInfo)
		}
	} else {
		return errno.VideoNotExit
	}
	return nil
}

func (*ActionService) CommentList(_ context.Context, req *service.DouyinCommentListRequest, resp *service.DouyinCommentListResponse) error {
	userId := utils.GetUserId(req.GetToken())
	var (
		commentInfos []*service.Comment
	)

	videoId := req.GetVideoId()
	svid := strconv.Itoa(int(videoId))
	if vexit := db.CheckVideoExit(videoId); vexit {
		if exit := redis.CheckVideoIdInComments(svid); exit {
			ids := redis.GetCommentIdsInComments(svid)
			for i := 0; i < len(ids); i++ {
				var commentInfo service.Comment
				var user service.User
				commentInfo.User = &user
				comment := db.GetCommentById(ids[i])
				// 用户信息
				userInfo := db.GetUserInfoById(comment.UserId)
				// 点赞评论信息
				followInfo := db.GetUserFollowInfo(comment.UserId, userId)
				commentInfo.Id = &comment.Id
				commentInfo.Content = &comment.CommentText
				format := comment.CommentTime.Format("01-02")
				commentInfo.CreateDate = &format
				commentInfo.User.Id = &userInfo.UserId
				commentInfo.User.Name = &userInfo.Name
				commentInfo.User.FollowCount = &followInfo.FollowCount
				commentInfo.User.FollowerCount = &followInfo.FollowerCount
				commentInfo.User.IsFollow = &followInfo.IsFollow
				commentInfos = append(commentInfos, &commentInfo)
			}
		} else {
			redis.AddVideoIdInComments(svid)
			redis.AddExpireInComments(svid)
			commentList := db.GetCommentList(videoId)
			for i := 0; i < len(commentList); i++ {
				var commentInfo service.Comment
				var user service.User
				commentInfo.User = &user
				if save := redis.AddCommentIdInComments(svid, commentList[i].Id); !save {
					redis.DeleteVideoIdInComments(svid)
				}
				// 用户信息
				userInfo := db.GetUserInfoById(commentList[i].UserId)
				// 点赞评论信息
				followInfo := db.GetUserFollowInfo(commentList[i].UserId, userId)
				commentInfo.Id = &commentList[i].Id
				commentInfo.Content = &commentList[i].CommentText
				format := commentList[i].CommentTime.Format("01-02")
				commentInfo.CreateDate = &format
				commentInfo.User.Id = &userInfo.UserId
				commentInfo.User.Name = &userInfo.Name
				commentInfo.User.FollowCount = &followInfo.FollowCount
				commentInfo.User.FollowerCount = &followInfo.FollowerCount
				commentInfo.User.IsFollow = &followInfo.IsFollow
				commentInfos = append(commentInfos, &commentInfo)
			}
		}
	} else {
		return errno.VideoNotExit
	}
	pack.BuildCommentListResp(resp, commentInfos)
	return nil
}
