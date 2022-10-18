package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/bellacbs/Delivery/simulator/application/route"
	"github.com/bellacbs/Delivery/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
//{"clientId":"1","routeId":"1"}
//{"clientId":"2","routeId":"2"}
//{"clientId":"3","routeId":"3"}
func Produce(message *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(message.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println("erro to export json positions", err.Error())
	}
	for _, position := range positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Microsecond * 500)
	}
}
