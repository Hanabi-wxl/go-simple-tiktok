package consts

import "time"

const (
	ETCDAddr               = "127.0.0.1:2379"
	ActionServiceName      = "rpcActionService"
	ActionServiceAddr      = "127.0.0.1:8081"
	MysqlDSN               = "root:@tcp(localhost:3306)/simple_tiktok?charset=utf8&parseTime=True&loc=Local"
	RedisHost              = "127.0.0.1:6379"
	RedisPassword          = "root"
	RabbitMQHost           = "amqp://guest:guest@127.0.0.1:5672/"
	RMQStartAddQueueName   = "star_add"
	RMQStartDelQueueName   = "star_del"
	RMQCommentDelQueueName = "comment_del"
)

var DefaultCode int32 = 0
var DefaultMsg = "success"
var DefaultRedisValue int64 = -1
var DefaultRedisTimeOut = time.Second * 60 * 60 * 24 * 30
