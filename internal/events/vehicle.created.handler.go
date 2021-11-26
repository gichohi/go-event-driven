package events

import (
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	KafkaClient "github.com/gichohi/go-cqrs-kafka-grpc/pkg/kafka"
	"log"
)

type VehicleCreatedEventHandler interface {
	Handle(event *VehicleCreatedEvent) error
}

type VehicleCreatedHandler struct {

}

func NewVehicleCreatedHandler() *VehicleCreatedHandler {
	return &VehicleCreatedHandler{}
}

func (vehicleCreatedHandler *VehicleCreatedHandler) Handle(event *VehicleCreatedEvent) {

	log.Println("Handling")
	vehicle := &models.Vehicle{VehicleID: event.CreateDto.VehicleID, Name: event.CreateDto.Name, Description: event.CreateDto.Description, Make: event.CreateDto.Make, Model: event.CreateDto.Model, Year: event.CreateDto.Year, Price: event.CreateDto.Price}
	topic := "VehicleCreated"

	KafkaClient.Publish(topic, vehicle)
}
