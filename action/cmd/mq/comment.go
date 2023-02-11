package mq

import (
	"action/cmd/dal/db"
	"action/cmd/model"
	"action/pkg/consts"
	"encoding/json"
	"github.com/streadway/amqp"
)

type CommentMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

var CommentDelQue *CommentMQ

func initComment() {
	CommentDelQue = newCommentMQ(consts.RMQCommentDelQueueName)
	go CommentDelQue.Consumer()
}

func newCommentMQ(queueName string) *CommentMQ {
	commentMQ := &CommentMQ{
		RabbitMQ:  *MQ,
		queueName: queueName,
	}
	cha, err := commentMQ.conn.Channel()
	commentMQ.channel = cha
	MQ.SendMqErr(err, "获取通道失败")
	return commentMQ
}

// Publish 消息发布
func (cmq *CommentMQ) Publish(message []byte) {
	_, err := cmq.channel.QueueDeclare(
		cmq.queueName,
		true,  //是否持久化
		false, //是否为自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		panic(err)
	}
	err1 := cmq.channel.Publish(
		cmq.exchange,
		cmq.queueName,
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
func (cmq *CommentMQ) Consumer() {
	_, err := cmq.channel.QueueDeclare(
		cmq.queueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	messages, err1 := cmq.channel.Consume(
		cmq.queueName,
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
	switch cmq.queueName {
	case consts.RMQCommentDelQueueName:
		//点赞消费队列
		go cmq.consumerCommentDel(messages)
	}
	<-forever
}

// consumerCommentDel 删除评论消费队列
func (cmq *CommentMQ) consumerCommentDel(messages <-chan amqp.Delivery) {
	for d := range messages {
		var comment model.MQComment
		err := json.Unmarshal(d.Body, &comment)
		if err != nil {
			panic(err)
		}
		db.DeleteComment(comment.CommentId)
		_ = d.Ack(false)
	}
}
