package kafka

import (
	"context"
	"encoding/json"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"log"
	"time"
)


func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  3,
		Compression:  compress.Snappy,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func Publish(topic string, vehicle *models.Vehicle){
	writer := newKafkaWriter("kafka:9092", topic)
	crateVehicleJson, _ := json.Marshal(vehicle)

	msg := kafka.Message{
		Value: crateVehicleJson,
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatalf("Producer Error: %s", err.Error())
	}
}