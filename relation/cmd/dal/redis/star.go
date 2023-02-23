package redis

import (
	"log"
	"relation/pkg/consts"
	"strconv"
)

// CheckUserIdExistInStar
// @Description: 检查是否存在userId
// @auth sinre 2023-02-10 21:04:53
// @param ctx 上下文
// @param uid 用户id
// @return bool 存在标志
func CheckUserIdExistInStar(suid string) bool {
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
	if n, err := RdStar.LPush(rdContext, suid, vid).Result(); n > 0 {
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	}
	return false
}

// CreateUserIdInStar
// @Description: 添加Key：userId
// @auth sinre 2023-02-10 22:56:00
// @param ctx 上下文
// @param uid userId
func CreateUserIdInStar(suid string) {
	if _, err := RdStar.LPush(rdContext, suid, consts.DefaultRedisValue).Result(); err != nil {
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
	videoIdList, err1 := RdStar.LRange(rdContext, suid, 0, -1).Result()
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
