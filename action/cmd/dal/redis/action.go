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
	videoIdList, err := RdStar.SMembers(rdContext, suid).Result()
	if err != nil {
		log.Println(err.Error())
	} else {
		for i := 0; i < len(videoIdList); i++ {
			vid, err := strconv.Atoi(videoIdList[i])
			if err != nil {
				ids = append(ids, int64(vid))
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
