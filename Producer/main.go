package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "test_topic"

	for _, word := range []string{"Welcome", "to", "Kafka", "with", "Golang"} {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)

		if err != nil {
			fmt.Printf("Failed to produce message: %s\n", err)
		}
	}

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
}
