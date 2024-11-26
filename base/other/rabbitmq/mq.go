package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const (
	queueName      = "task_queue"
	producers      = 10000           // 生产者数量
	messagesPerSec = 10              // 每秒产生的消息数量
	runDuration    = 1 * time.Minute // 运行持续时间
)

func main() {
	conn, err := amqp.Dial("amqp://admin:antiy123.@10.251.36.13:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// 创建队列
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName, // 队列名
		true,      // 持久化
		false,     // 仅用于消费，生产者不删除
		false,     // 排他性
		false,     // 自动删除
		nil,       // 额外属性
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// 启动消费者
	go startConsumer(conn)

	// 启动生产者
	var wg sync.WaitGroup
	for i := 0; i < producers; i++ {
		wg.Add(1)
		go startProducer(ch, &wg, i)
	}

	// 运行 1 分钟后停止
	time.Sleep(runDuration)
	fmt.Println("Stopping producers...")
	ch.Close()
	wg.Wait()
	fmt.Println("All producers have finished.")
}

func startProducer(ch *amqp.Channel, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second / 10) // 每秒发送 10 条消息
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for j := 0; j < messagesPerSec; j++ {
				message := fmt.Sprintf("Hello from producer %d", id)
				err := ch.Publish(
					"",        // 默认交换机
					queueName, // 队列名
					false,     // 必须
					false,     // 立即
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(message),
					},
				)
				if err != nil {
					log.Printf("Failed to publish a message: %s", err)
				} else {
					log.Printf("Producer %d sent: %s", id, message)
				}
			}
		}
	}
}

func startConsumer(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",    // 消费者唯一标识
		true,  // 自动确认
		false, // 排他性
		false, // 不等待
		false, // 不自动关闭
		nil,   // 额外属性
	)
	if err != nil {
		log.Fatalf("Failed to register consumption: %s", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Consumer received: %s", msg.Body)
		}
	}()

	// Wait to exit
	select {}
}
