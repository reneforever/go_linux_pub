package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 专门写入kafka模块

var (
	client sarama.SyncProducer //声明一个全局连接kafka的client
)

// 初始化client
func Init(addrs []string) (err error) {
	// 新建config
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
		return
	}
	return
}

// 数据建立，格式化，发送
func SendToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送到kafka
	var (
		pid    int32
		offset int64
		err    error
	)
	pid, offset, err = client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err: ", err)
		return
	}
	fmt.Printf("[pid: %v, offset: %v]\n", pid, offset)
}
