package mq

import (
	"action/pkg/consts"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn *amqp.Connection
	url  string
}

var MQ *RabbitMQ

func Init() {
	MQ = &RabbitMQ{
		url: consts.RabbitMQHost,
	}
	dial, err := amqp.Dial(MQ.url)
	if err != nil {
		MQ.SendMqErr(err, "mq连接失败")
	}
	MQ.conn = dial
	initStar()
	initComment()
}

func (r *RabbitMQ) SendMqErr(err error, message string) {
	if err != nil {
		log.Println(err, message)
		panic(err)
	}
}

func (r *RabbitMQ) Close() {
	_ = r.conn.Close()
}
