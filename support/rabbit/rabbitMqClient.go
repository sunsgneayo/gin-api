package client

import (
	"dpj-admin-api/config"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

const DefaultRetryInterval = 5 * time.Second

// RabbitMQ 结构体
type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Host      string
	Username  string
	Password  string
	Port      string
}

// ConnectionConfig 结构体用于动态设置连接配置
type ConnectionConfig struct {
	Host             string
	Username         string
	Password         string
	Port             string
	RetryInterval    time.Duration
	MaxRetryAttempts int
}

// defaultConnectionConfig 默认连接配置
var defaultConnectionConfig = &ConnectionConfig{
	Host:             config.Get("rabbit.host"),
	Username:         config.Get("rabbit.username"),
	Password:         config.Get("rabbit.password"),
	Port:             config.Get("rabbit.port"),
	RetryInterval:    DefaultRetryInterval,
	MaxRetryAttempts: 3,
}

// SetConnectionConfig 设置连接配置
func SetConnectionConfig(config *ConnectionConfig) {
	defaultConnectionConfig = config
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, args ...string) (*RabbitMQ, error) {
	// 使用默认连接配置
	var exchange, key string
	// Check if exchange and key arguments are provided
	if len(args) >= 1 {
		exchange = args[0]
	}
	if len(args) >= 2 {
		key = args[1]
	}
	return newRabbitMQ(queueName, exchange, key, defaultConnectionConfig)
}

// newRabbitMQ 使用给定的连接配置创建结构体实例
func newRabbitMQ(queueName, exchange, key string, config *ConnectionConfig) (*RabbitMQ, error) {
	var conn *amqp.Connection
	var channel *amqp.Channel

	port, _ := strconv.Atoi(config.Port)
	mqURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, port)

	// 尝试建立连接，支持重试
	for attempt := 1; attempt <= config.MaxRetryAttempts; attempt++ {
		var err error
		conn, err = amqp.Dial(mqURL)
		if err == nil {
			break
		}

		log.Printf("Failed to connect rabbitmq (Attempt %d/%d): %s", attempt, config.MaxRetryAttempts, err)
		time.Sleep(config.RetryInterval)
	}

	if conn == nil {
		return nil, fmt.Errorf("failed to connect rabbitmq after %d attempts", config.MaxRetryAttempts)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %s", err)
	}

	return &RabbitMQ{
		Connection: conn,
		Channel:    channel,
		QueueName:  queueName,
		Exchange:   exchange,
		Key:        key,
		Host:       config.Host,
		Username:   config.Username,
		Password:   config.Password,
		Port:       config.Port,
	}, nil
}

// Destroy 断开Connection和Channel
func (r *RabbitMQ) Destroy() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
}

// failOnErr 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) error {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		return fmt.Errorf("%s:%s", message, err)
	}
	return nil
}

// setConnectionConfig 设置连接配置
func (r *RabbitMQ) setConnectionConfig(config *ConnectionConfig) {
	defaultConnectionConfig = config
}

// NewRabbitMQSimple 简单模式下创建RabbitMQ实例
func NewRabbitMQSimple(queueName string) (*RabbitMQ, error) {
	return newRabbitMQ(queueName, "", "", defaultConnectionConfig)
}

// PublishSimple 直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	err = r.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "Failed to publish message to queue")
}

// ConsumeSimple Simple模式下消费者
func (r *RabbitMQ) ConsumeSimple(consumeLogic func(msg amqp.Delivery)) {
	q, err := r.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	msgs, err := r.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			consumeLogic(d)
		}
	}()

	log.Printf("队列： [%s] 消费进程已开启. 要退出，请按 CTRL+C ", r.QueueName)
	<-forever
}
