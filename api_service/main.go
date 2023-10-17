package main

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/golang/protobuf/proto"
    "myservice" // Import the generated protobuf code
)

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
    if err != nil {
        panic(err)
    }
    defer p.Close()

    // Create a protobuf message
    msg := &myservice.Message{Content: "Hello, Kafka!"}
    data, err := proto.Marshal(msg)
    if err != nil {
        panic(err)
    }

    // Produce the message to a Kafka topic
    err = p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:         data,
    }, nil)
    if err != nil {
        panic(err)
    }

    p.Flush(15 * 1000)
}

