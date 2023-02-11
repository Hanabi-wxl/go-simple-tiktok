package mq

import (
	"action/cmd/dal/db"
	"action/cmd/model"
	"action/pkg/consts"
	"encoding/json"
	"github.com/streadway/amqp"
)

type StarMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

var StarAddQue *StarMQ
var StarDelQue *StarMQ

func initStar() {
	StarAddQue = newStarMQ(consts.RMQStartAddQueueName)
	go StarAddQue.Consumer()

	StarDelQue = newStarMQ(consts.RMQStartDelQueueName)
	go StarDelQue.Consumer()
}

func newStarMQ(queueName string) *StarMQ {
	starMQ := &StarMQ{
		RabbitMQ:  *MQ,
		queueName: queueName,
	}
	cha, err := starMQ.conn.Channel()
	starMQ.channel = cha
	MQ.SendMqErr(err, "获取通道失败")
	return starMQ
}

// Publish 消息发布
func (smq *StarMQ) Publish(message []byte) {
	_, err := smq.channel.QueueDeclare(
		smq.queueName,
		true,  //是否持久化
		false, //是否为自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		panic(err)
	}
	err1 := smq.channel.Publish(
		smq.exchange,
		smq.queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         message,
		})
	if err1 != nil {
		panic(err)
	}
}

// Consumer 消息消费
func (smq *StarMQ) Consumer() {
	_, err := smq.channel.QueueDeclare(
		smq.queueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	messages, err1 := smq.channel.Consume(
		smq.queueName,
		"",    //用来区分多个消费者
		false, //是否自动应答
		false, //是否具有排他性
		false, // rabbitmq不支持
		false, //消息队列是否阻塞
		nil,
	)
	if err1 != nil {
		panic(err1)
	}
	forever := make(chan bool)
	switch smq.queueName {
	case consts.RMQStartAddQueueName:
		//点赞消费队列
		go smq.consumerStarAdd(messages)
	case consts.RMQStartDelQueueName:
		//取消赞消费队列
		go smq.consumerStarDel(messages)
	}
	<-forever
}

// consumerStarAdd 点赞消费队列
func (smq *StarMQ) consumerStarAdd(messages <-chan amqp.Delivery) {
	for d := range messages {
		var ss model.MQStar
		err := json.Unmarshal(d.Body, &ss)
		if err != nil {
			panic(err)
		}
		db.CreateFavorite(ss.VideoId, ss.UserId)
		// 手动确认消息已消费
		_ = d.Ack(false)
	}
}

// consumerStarDel 取消点赞消费队列
func (smq *StarMQ) consumerStarDel(messages <-chan amqp.Delivery) {
	for d := range messages {
		var ss model.MQStar
		err := json.Unmarshal(d.Body, &ss)
		if err != nil {
			panic(err)
		}
		db.DeleteFavorite(ss.VideoId, ss.UserId)
		_ = d.Ack(false)
	}
}
