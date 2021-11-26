package kafka

import (
	"context"
	"encoding/json"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/commands"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/dto"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

type Consumer interface {
	Read(topic string) error
}

type consumer struct {
	log           logger.Logger
	cfg           *config.Config
}


func getKafkaReader(kafkaURL string, topic string, groupId string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func CreateVehicleReader(topic string){
	kafkaURL := "kafka:9092"
	groupId := "vehicle-group"
	reader := getKafkaReader(kafkaURL, topic, groupId)

	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {

		}
	}(reader)

	message, _ := reader.ReadMessage(context.Background())
	log.Printf("Message Received: %s\n", message.Value)

	d := &dto.VehicleDto{}
	err := json.Unmarshal(message.Value, d)
	if err != nil {
		return
	}
	vehicleDto := dto.VehicleDto{VehicleID: d.VehicleID,Name: d.Name,Description: d.Description, Make: d.Make, Model: d.Model,
	Year: d.Year, Price: d.Price}
	commands.NewCreateVehicleCommand(&vehicleDto)
}

