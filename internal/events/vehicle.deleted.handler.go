package events

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	KafkaClient "github.com/gichohi/go-cqrs-kafka-grpc/pkg/kafka"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
)

type VehicleDeletedEventHandler interface {
	Handle(ctx context.Context, event *DeleteVehicleEvent) error
}

type VehicleDeletedHandler struct {
	log           logger.Logger
	cfg           *config.Config
}

func NewDeletedVehicleHandler(log logger.Logger, cfg *config.Config) *VehicleDeletedHandler {
	return &VehicleDeletedHandler{log: log, cfg: cfg}
}

func (vehicleDeletedHandler *VehicleDeletedHandler) Handle(ctx context.Context, event *DeleteVehicleEvent) {

	vehicle := &models.Vehicle{VehicleID: event.VehicleID}

	topic := "Vehicle Created"

	KafkaClient.Publish(topic, vehicle)
}
