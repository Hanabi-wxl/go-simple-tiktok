package redis

import (
	"log"
	"relation/pkg/consts"
	"strconv"
)

// CheckVideoIdExistInStars
// @Description: 检查是否存在Key: videoId
// @auth sinre 2023-02-10 23:31:53
// @param ctx 上下文
// @param svid videoId
// @return bool 存在标志
func CheckVideoIdExistInStars(svid string) bool {
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
	if _, err := RdStars.LPush(rdContext, svid, suid).Result(); err != nil {
		log.Println(err.Error())
	}
	return true
}

// CreateVideoIdInStars
// @Description: 创建Key: videoId
// @auth sinre 2023-02-10 23:36:16
// @param ctx 上下文
// @param svid 视频id
func CreateVideoIdInStars(svid string) {
	if _, err := RdStars.LPush(rdContext, svid, consts.DefaultRedisValue).Result(); err != nil {
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
	userList, err1 := RdStars.LRange(rdContext, svid, 0, -1).Result()
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
