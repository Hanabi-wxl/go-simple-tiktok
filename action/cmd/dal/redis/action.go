package redis

import (
	"action/pkg/consts"
	"log"
	"strconv"
)

// CheckUserIdExitInStar
// @Description: 检查是否存在userId
// @auth sinre 2023-02-10 21:04:53
// @param ctx 上下文
// @param uid 用户id
// @return bool 存在标志
func CheckUserIdExitInStar(suid string) bool {
	if n, err := RdStar.Exists(rdContext, suid).Result(); n > 0 {
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	}
	return false
}

// AddVideoIdInStar
// @Description: 创建点赞信息
// @auth sinre 2023-02-10 22:48:34
// @param ctx 上下文
// @param uid 用户id
// @param vid 视频id
// @return bool 保存成功
func AddVideoIdInStar(suid string, vid int64) bool {
	if _, err := RdStar.SAdd(rdContext, suid, vid).Result(); err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// RemoveVideoIdInStar
// @Description: 删除点赞信息
// @auth sinre 2023-02-10 23:55:40
// @param ctx 上下文
// @param uid 用户id
// @param vid 视频id
func RemoveVideoIdInStar(suid string, vid int64) {
	if _, err := RdStar.SRem(rdContext, suid, vid).Result(); err != nil {
		log.Println(err.Error())
	}
}

// CreateUserIdInStar
// @Description: 添加Key：userId
// @auth sinre 2023-02-10 22:56:00
// @param ctx 上下文
// @param uid userId
func CreateUserIdInStar(suid string) {
	if _, err := RdStar.SAdd(rdContext, suid, consts.DefaultRedisValue).Result(); err != nil {
		log.Println(err.Error())
	}
}

// DeleteUserIdInStar
// @Description: 删除Key: userId
// @auth sinre 2023-02-10 23:27:25
// @param ctx 上下文
// @param suid 用户id
func DeleteUserIdInStar(suid string) {
	_, _ = RdStar.Del(rdContext, suid).Result()
}

// AddExpireInStar
// @Description: 添加过期时间
// @auth sinre 2023-02-10 23:28:09
// @param ctx 上下文
// @param suid 用户id
func AddExpireInStar(suid string) {
	if _, err := RdStar.Expire(rdContext, suid, consts.DefaultRedisTimeOut).Result(); err != nil {
		log.Println(err.Error())
	}
}

// GetVideoIdsInStar
// @Description: 获取用户所有点赞的视频id
// @auth sinre 2023-02-11 13:54:20
// @param suid 用户id
// @return ids 视频id
func GetVideoIdsInStar(suid string) (ids []int64) {
	videoIdList, err1 := RdStar.SMembers(rdContext, suid).Result()
	if err1 != nil {
		log.Println(err1.Error())
	} else {
		for i := 0; i < len(videoIdList); i++ {
			vid, err := strconv.Atoi(videoIdList[i])
			vvid := int64(vid)
			if err == nil && vvid != consts.DefaultRedisValue {
				ids = append(ids, vvid)
			}
		}
	}
	return ids
}

// CheckVideoIdExitInStars
// @Description: 检查是否存在Key: videoId
// @auth sinre 2023-02-10 23:31:53
// @param ctx 上下文
// @param svid videoId
// @return bool 存在标志
func CheckVideoIdExitInStars(svid string) bool {
	if n, err := RdStars.Exists(rdContext, svid).Result(); n > 0 {
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	}
	return false
}

// AddUserIdInStars
// @Description: 创建videoId对于的userId
// @auth sinre 2023-02-10 23:35:19
// @param ctx 上下文
// @param svid 视频id
// @param suid 用户id
func AddUserIdInStars(svid string, suid int64) bool {
	if _, err := RdStars.SAdd(rdContext, svid, suid).Result(); err != nil {
		log.Println(err.Error())
	}
	return true
}

// RemoveUserIdInStars
// @Description: 删除点赞的用户
// @auth sinre 2023-02-11 00:02:30
// @param ctx 上下文
// @param svid 视频id
// @param suid 用户id
func RemoveUserIdInStars(svid string, suid int64) {
	if _, err := RdStars.SRem(rdContext, svid, suid).Result(); err != nil {
		log.Println(err.Error())
	}
}

// CreateVideoIdInStars
// @Description: 创建Key: videoId
// @auth sinre 2023-02-10 23:36:16
// @param ctx 上下文
// @param svid 视频id
func CreateVideoIdInStars(svid string) {
	if _, err := RdStars.SAdd(rdContext, svid, consts.DefaultRedisValue).Result(); err != nil {
		log.Println(err.Error())
	}
}

// DeleteVideoIdInStars
// @Description: 删除Key: videoId
// @auth sinre 2023-02-10 23:36:49
// @param ctx 上下文
// @param svid 视频id
func DeleteVideoIdInStars(svid string) {
	_, _ = RdStars.Del(rdContext, svid).Result()
}

// AddExpireInStars
// @Description: 添加过期时间
// @auth sinre 2023-02-10 23:37:17
// @param ctx 上下文
// @param svid 视频id
func AddExpireInStars(svid string) {
	if _, err := RdStars.Expire(rdContext, svid, consts.DefaultRedisTimeOut).Result(); err != nil {
		log.Println(err.Error())
	}
}

// GetUserIdsInStars
// @Description: 获取视频点赞用户的列表
// @auth sinre 2023-02-11 15:35:20
// @param svid 视频id
// @return ids 用户id列表
func GetUserIdsInStars(svid string) (ids []int64) {
	userList, err1 := RdStars.SMembers(rdContext, svid).Result()
	if err1 != nil {
		log.Println(err1.Error())
	} else {
		for i := 0; i < len(userList); i++ {
			uid, err := strconv.Atoi(userList[i])
			uuid := int64(uid)
			if err == nil && uuid != consts.DefaultRedisValue {
				ids = append(ids, uuid)
			}
		}
	}
	return ids
}

// CheckVideoIdInComments
// @Description: 查看是否存在key: videoId
// @auth sinre 2023-02-11 15:02:53
// @param vid 视频id
// @return bool 存在标志
func CheckVideoIdInComments(vid string) bool {
	if n, err := RdComments.Exists(rdContext, vid).Result(); n > 0 {
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	}
	return false
}

// AddVideoIdInComments
// @Description: 添加key:videoId
// @auth sinre 2023-02-11 15:03:29
// @param vid 视频id
func AddVideoIdInComments(vid string) {
	if _, err := RdComments.SAdd(rdContext, vid, consts.DefaultRedisValue).Result(); err != nil {
		log.Println(err.Error())
	}
}

// AddExpireInComments
// @Description: 添加过期时间
// @auth sinre 2023-02-11 15:05:51
// @param vid 视频id
func AddExpireInComments(vid string) {
	if _, err := RdComments.Expire(rdContext, vid, consts.DefaultRedisTimeOut).Result(); err != nil {
		log.Println(err.Error())
	}
}

// AddCommentIdInComments
// @Description: 添加评论id
// @auth sinre 2023-02-11 15:13:10
// @param svid 视频id
// @param cid 评论id
// @return bool 添加成功
func AddCommentIdInComments(svid string, cid int64) bool {
	if _, err := RdComments.SAdd(rdContext, svid, cid).Result(); err != nil {
		log.Println(err.Error())
		return false
	} else {
		return true
	}
}

// DeleteVideoIdInComments
// @Description: 删除key: videoId
// @auth sinre 2023-02-11 15:15:03
// @param svid 视频id
func DeleteVideoIdInComments(svid string) {
	if _, err := RdComments.Del(rdContext, svid).Result(); err != nil {
		log.Println(err.Error())
	}
}

// RemoveCommentIdInComments
// @Description: 删除key: videoId
// @auth sinre 2023-02-11 15:15:03
// @param svid 视频id
func RemoveCommentIdInComments(svid string, cid int64) {
	if _, err := RdComments.SRem(rdContext, svid, cid).Result(); err != nil {
		log.Println(err.Error())
	}
}

// GetCommentIdsInComments
// @Description: 获取视频的所有评论id
// @auth sinre 2023-02-11 15:40:43
// @param svid 视频id
// @return ids 评论列表
func GetCommentIdsInComments(svid string) (ids []int64) {
	cidl, err1 := RdComments.SMembers(rdContext, svid).Result()
	if err1 != nil {
		log.Println(err1.Error())
	} else {
		for i := 0; i < len(cidl); i++ {
			cid, err := strconv.Atoi(cidl[i])
			ccid := int64(cid)
			if err == nil && ccid != consts.DefaultRedisValue {
				ids = append(ids, ccid)
			}
		}
	}
	return ids
}
