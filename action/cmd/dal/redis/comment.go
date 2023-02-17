package redis

import (
	"action/pkg/consts"
	"log"
	"strconv"
)

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
	if _, err := RdComments.RPush(rdContext, vid, consts.DefaultRedisValue).Result(); err != nil {
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
	if _, err := RdComments.RPush(rdContext, svid, cid).Result(); err != nil {
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
	if _, err := RdComments.LRem(rdContext, svid, 1, cid).Result(); err != nil {
		log.Println(err.Error())
	}
}

// GetCommentIdsInComments
// @Description: 获取视频的所有评论id
// @auth sinre 2023-02-11 15:40:43
// @param svid 视频id
// @return ids 评论列表
func GetCommentIdsInComments(svid string) (ids []int64) {
	cidl, err1 := RdComments.LRange(rdContext, svid, 0, -1).Result()
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
