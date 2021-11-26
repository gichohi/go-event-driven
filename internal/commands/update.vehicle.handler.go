package commands

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/repository"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"log"
)


type UpdateVehicleCmdHandler interface {
	Handle(ctx context.Context, command *UpdateVehicleCommand)
}

type UpdateVehicleHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
}

func NewUpdateVehicleHandler(log logger.Logger, cfg *config.Config) *UpdateVehicleHandler {
	return &UpdateVehicleHandler{log: log, cfg: cfg}
}

func (updateVehicleHandler *UpdateVehicleHandler) Handle(ctx context.Context, command *UpdateVehicleCommand) {

	vehicleDto := &models.Vehicle{VehicleID: command.VehicleID, Name: command.Name, Description: command.Description, Make: command.Make, Model: command.Model, Year: command.Year, Price: command.Price}

	_, err := updateVehicleHandler.pgRepo.UpdateVehicle(ctx, vehicleDto)
	if err != nil {
		log.Fatal(err)
	}
}
