package redis

import (
	"context"
	"core/pkg/consts"
	"github.com/go-redis/redis/v8"
)

var rdContext = context.Background()

var RdStars *redis.Client
var RdStar *redis.Client
var RdWorks *redis.Client
var RdComments *redis.Client

func Init() {
	// 点赞列表 - db0 - star
	// 获赞数量 - db1 - stars
	// 发布列表 - db2 - works

	// 保存视频被哪些用户点赞 videoId: userId
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

	RdWorks = redis.NewClient(&redis.Options{
		Addr:     consts.RedisHost,
		Password: consts.RedisPassword,
		DB:       2,
	})

	// 保存视频的评论信息 videoId: commentId
	RdComments = redis.NewClient(&redis.Options{
		Addr:     consts.RedisHost,
		Password: consts.RedisPassword,
		DB:       3,
	})
}
