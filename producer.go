package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
)

func main() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer error", err)
	}
	// publish without goroutene
	message := &sarama.ProducerMessage{
		Topic: "sampleTopic", Value: sarama.StringEncoder("golang message"),
	}
	p, o, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}
	fmt.Println("partition", p)
	fmt.Println("offset", o)

}
