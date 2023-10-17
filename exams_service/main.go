package main

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/golang/protobuf/proto"
    "myservice" // Import the generated protobuf code
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "my-group",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    // Subscribe to a Kafka topic
    err = c.SubscribeTopics([]string{"my-topic"}, nil)
    if err != nil {
        panic(err)
    }

    for {
        msg, err := c.ReadMessage(-1)
        if err == nil {
            // Parse the protobuf message
            var receivedMsg myservice.Message
            err = proto.Unmarshal(msg.Value, &receivedMsg)
            if err != nil {
                panic(err)
            }
            // Process the message
            println("Received message:", receivedMsg.Content)
        }
    }
}

