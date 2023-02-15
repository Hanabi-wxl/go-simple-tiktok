package consts

import "time"

const (
	//ETCDAddr               = "192.168.64.3:2379"
	//ActionServiceAddr      = "0.0.0.0:8081"
	//RedisHost              = "192.168.64.5:6379"
	//MysqlDSN               = "root:@tcp(192.168.64.2:3306)/simple_tiktok?charset=utf8&parseTime=True&loc=Local"
	//RabbitMQHost           = "amqp://guest:guest@192.168.64.4:5672/"
	//BackgroundImgUrl       = "http://tiktok.sinre.top/static/avatar/back.jpg"

	ETCDAddr          = "127.0.0.1:2379"
	ActionServiceName = "rpcActionService"
	ActionServiceAddr = "127.0.0.1:8081"
	MysqlDSN          = "root:@tcp(localhost:3306)/simple_tiktok?charset=utf8&parseTime=True&loc=Local"
	BackgroundImgUrl  = "http://192.168.1.8:8000/static/avatar/back.jpg"
	RedisHost         = "127.0.0.1:6379"
	RabbitMQHost      = "amqp://guest:guest@127.0.0.1:5672/"

	RedisPassword          = "root"
	RMQStartAddQueueName   = "star_add"
	RMQStartDelQueueName   = "star_del"
	RMQCommentDelQueueName = "comment_del"
)

var DefaultCode int32 = 0
var DefaultMsg = "success"
var DefaultRedisValue int64 = -1
var DefaultRedisTimeOut = time.Second * 60 * 60 * 24 * 30
