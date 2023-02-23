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
	"sync"
)

func (*ActionService) FavoriteAction(_ context.Context, req *service.DouyinFavoriteActionRequest, resp *service.DouyinFavoriteActionResponse) error {
	acType := req.GetActionType()
	token := req.GetToken()
	videoId := req.GetVideoId()
	uid := utils.GetUserId(token)
	suid := strconv.Itoa(int(uid))
	svid := strconv.Itoa(int(videoId))
	if exist := db.CheckVideoExist(videoId); exist {
		body, _ := json.Marshal(model.MQStar{UserId: uid, VideoId: videoId})
		if acType == 1 {
			// 维护redis数据库: Star
			// 查询是否存在userId
			if exist1 := redis.CheckUserIdExistInStar(suid); exist1 {
				// 存在Key: userId则创建点赞信息至redis
				if save := redis.AddVideoIdInStar(suid, videoId); save {
					// 使用mq同步信息至db
					mq.StarAddQue.Publish(body)
				} else {
					return errno.RepeatStar
				}
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
			if exist1 := redis.CheckVideoIdExistInStars(svid); exist1 {
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
			if exist1 := redis.CheckUserIdExistInStar(suid); exist1 {
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

			if exist1 := redis.CheckVideoIdExistInStars(svid); exist1 {
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
		return errno.VideoNotExist
	}
	pack.BuildFavoriteActionResp(resp)
	return nil
}

func (*ActionService) FavoriteList(_ context.Context, req *service.DouyinFavoriteListRequest, resp *service.DouyinFavoriteListResponse) error {
	var userId int64
	var fids []int64
	token := req.GetToken()
	checkId := req.GetUserId()
	if existUser := db.CheckUserIdExist(checkId); !existUser {
		return errno.UserNotExistErr
	}
	userId = utils.GetUserId(token)
	suid := strconv.Itoa(int(userId))
	scid := strconv.Itoa(int(checkId))
	// 作者信息
	infoById := db.GetUserInfoById(checkId)
	// 查询自己时userId == checkId
	followInfo := db.GetUserFollowInfo(checkId, userId)
	totalFav, workCount, starCount := getUserCountInfo(userId)
	author := pack.BuildAuthor(infoById, followInfo, checkId, totalFav, workCount, starCount)
	if exist := redis.CheckUserIdExistInStar(scid); exist {
		// 查询视频作者喜欢列表
		fids = redis.GetVideoIdsInStar(scid)
	} else {
		// 不存在Key则新增Key
		redis.CreateUserIdInStar(suid)
		// 添加过期时间
		redis.AddExpireInStar(suid)
		// 获取用户所有点赞信息 保存至redis
		favoriteList := db.GetFavoriteListByUserId(checkId)
		if len(favoriteList) > 0 {
			for _, likeVideoId := range favoriteList {
				if save := redis.AddVideoIdInStar(suid, likeVideoId.VideoId); !save {
					// 如果redis保存出现问题则删掉该Key
					redis.DeleteUserIdInStar(suid)
				}
			}
			for i := 0; i < len(favoriteList); i++ {
				fids = append(fids, favoriteList[i].VideoId)
			}
		}
	}
	favIdLen := len(fids)
	var wg sync.WaitGroup
	wg.Add(favIdLen)
	videoInfos := make([]*service.Video, favIdLen)
	for i, fid := range fids {
		// 查询视频信息
		go addToFavList(i, fid, userId, author, &videoInfos, &wg)
	}
	wg.Wait()
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
		followInfo model.FollowInfo
	)

	token := req.GetToken()
	userId = utils.GetUserId(token)
	if exist := db.CheckVideoExist(videoId); exist {
		if acType == 1 {
			comment := db.CreateComment(userId, videoId, req.GetCommentText())
			// 判断是否存在key: videoId
			if exist1 := redis.CheckVideoIdInComments(svid); exist1 {
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
			userInfo := db.GetUserInfoById(userId)
			followInfo = db.GetUserFollowInfo(userId, userId)
			totalFav, workCount, starCount := getUserCountInfo(userId)
			author := pack.BuildAuthor(userInfo, followInfo, userId, totalFav, workCount, starCount)
			pack.BuildCommentActionResp(resp, &comment, author)
		} else if acType == 2 {
			if existc := db.CheckCommentExist(commentId); existc {
				// 判断是否有权删除评论
				if userId != db.GetCommentById(commentId).UserId {
					return errno.CommentDelErr
				}
				// mq信息序列化
				body, _ := json.Marshal(&model.MQComment{CommentId: commentId})
				// 判断是否存在key: videoIds
				if exist1 := redis.CheckVideoIdInComments(svid); exist1 {
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
				pack.BuildCommentActionResp(resp, nil, nil)
			} else {
				return errno.CommentNotExistErr
			}
		}
	} else {
		return errno.VideoNotExist
	}
	return nil
}

func (*ActionService) CommentList(_ context.Context, req *service.DouyinCommentListRequest, resp *service.DouyinCommentListResponse) error {
	userId := utils.GetUserId(req.GetToken())
	var ids []int64
	videoId := req.GetVideoId()
	svid := strconv.Itoa(int(videoId))
	if vexist := db.CheckVideoExist(videoId); vexist {
		if exist := redis.CheckVideoIdInComments(svid); exist {
			ids = redis.GetCommentIdsInComments(svid)
		} else {
			redis.AddVideoIdInComments(svid)
			redis.AddExpireInComments(svid)
			commentList := db.GetCommentList(videoId)
			for _, comment := range commentList {
				if save := redis.AddCommentIdInComments(svid, comment.Id); !save {
					redis.DeleteVideoIdInComments(svid)
				}
				ids = append(ids, comment.Id)
			}
		}
		var wg sync.WaitGroup
		comIdLen := len(ids)
		wg.Add(comIdLen)
		commentInfos := make([]*service.Comment, comIdLen)
		for i := 0; i < len(ids); i++ {
			go addToCommentList(i, ids[i], userId, &commentInfos, &wg)
		}
		wg.Wait()
		pack.BuildCommentListResp(resp, commentInfos)
	} else {
		return errno.VideoNotExist
	}
	return nil
}
