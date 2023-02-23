package redis

import (
	"action/pkg/consts"
	"log"
	"strconv"
)

func CheckUserIdExistInWorks(suid string) bool {
	if n, err := RdWorks.Exists(rdContext, suid).Result(); n > 0 {
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	}
	return false
}

func GetVideoIdsInWorks(suid string) (ids []int64) {
	videoIdList, err1 := RdWorks.LRange(rdContext, suid, 0, -1).Result()
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

func CreateUserIdInWorks(svid string) {
	if _, err := RdWorks.LPush(rdContext, svid, consts.DefaultRedisValue).Result(); err != nil {
		log.Println(err.Error())
	}
}

func AddExpireInWorks(svid string) {
	if _, err := RdWorks.Expire(rdContext, svid, consts.DefaultRedisTimeOut).Result(); err != nil {
		log.Println(err.Error())
	}
}

func AddVideoIdInWorks(suid string, svid int64) bool {
	if _, err := RdWorks.LPush(rdContext, suid, svid).Result(); err != nil {
		log.Println(err.Error())
		return false
	} else {
		return true
	}
}

func DeleteUserIdInWorks(suid string) {
	if _, err := RdWorks.Del(rdContext, suid).Result(); err != nil {
		log.Println(err.Error())
	}
}
