package redis

import (
	"action/pkg/consts"
	"context"
	"github.com/go-redis/redis/v8"
)

var rdContext = context.Background()

var RdStar *redis.Client
var RdStars *redis.Client

func Init() {
	// 点赞操作的主体为: 点赞者与视频
	// 保存用户点赞的视频 userId: videoId
	RdStar = redis.NewClient(&redis.Options{
		Addr:     consts.RedisHost,
		Password: consts.RedisPassword,
		DB:       0,
	})
	// 保存视频被哪些用户点赞 videoId: userId
	RdStars = redis.NewClient(&redis.Options{
		Addr:     consts.RedisHost,
		Password: consts.RedisPassword,
		DB:       1,
	})
}
