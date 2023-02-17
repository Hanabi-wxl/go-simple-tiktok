package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"relation/pkg/consts"
)

var rdContext = context.Background()

var RdStar *redis.Client
var RdStars *redis.Client

var RdComments *redis.Client

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
	// 保存视频的评论信息 videoId: commentId
	RdComments = redis.NewClient(&redis.Options{
		Addr:     consts.RedisHost,
		Password: consts.RedisPassword,
		DB:       2,
	})
}
