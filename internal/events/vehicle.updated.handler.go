package events

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	KafkaClient "github.com/gichohi/go-cqrs-kafka-grpc/pkg/kafka"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
)


type VehicleUpdatedEventHandler interface {
	Handle(ctx context.Context, event *UpdateVehicleEvent) error
}

type VehicleUpdatedHandler struct {
	log           logger.Logger
	cfg           *config.Config
}

func NewVehicleUpdatedHandler(log logger.Logger, cfg *config.Config) *VehicleUpdatedHandler {
	return &VehicleUpdatedHandler{log: log, cfg: cfg}
}

func (vehicleUpdatedHandler *VehicleUpdatedHandler) Handle(ctx context.Context, event *UpdateVehicleEvent) {
	vehicle := &models.Vehicle{VehicleID: event.VehicleID, Name: event.Name, Description: event.Description, Make: event.Make, Model: event.Model, Year: event.Year, Price: event.Price}
	topic := "Vehicle Created"

	KafkaClient.Publish(topic, vehicle)
}
