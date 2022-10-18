package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	kafkaProduce "github.com/bellacbs/Delivery/simulator/application/kafka"
	"github.com/bellacbs/Delivery/simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	messageChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(messageChannel)
	go consumer.Consume()
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "route.new-direction", producer)
	for message := range messageChannel {
		fmt.Println(string(message.Value))
		go kafkaProduce.Produce(message)
	}
}
