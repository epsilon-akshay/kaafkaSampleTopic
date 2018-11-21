package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, config)
	initialOffset := sarama.OffsetOldest
	pc, _ := consumer.ConsumePartition("sampleTopic", 0, initialOffset)
	for message := range pc.Messages() {
		fmt.Println("the message is", string(message.Value))
	}
}
