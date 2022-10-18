package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChannel chan *ckafka.Message
}

func NewKafkaConsumer(messageChannel chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChannel: messageChannel,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	consumer.SubscribeTopics(topics, nil)
	log.Println("Kafka consumer has been started")
	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			k.MessageChannel <- message
		}
	}

}
