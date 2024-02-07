package initialize

import (
	client "dpj-admin-api/support/rabbit"
	"github.com/streadway/amqp"
	"log"
)

func RabbitConsume() {
	go func() {

		rabbitmq, _ := client.NewRabbitMQ("queueName")

		rabbitmq.ConsumeSimple(func(msg amqp.Delivery) {
			log.Printf("接收到消费数据: %s", msg.Body)
		})
	}()
}
