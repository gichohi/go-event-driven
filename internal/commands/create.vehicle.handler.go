package commands

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/repository"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"log"
)

type CreateVehicleCmdHandler interface {
	Handle(ctx context.Context, command *CreateVehicleCommand)
}

type CreateVehicleHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
}

func NewCreateVehicleHandler(log logger.Logger, cfg *config.Config) *CreateVehicleHandler {
	return &CreateVehicleHandler{log: log, cfg: cfg}
}

func (createVehicleHandler *CreateVehicleHandler) Handle(ctx context.Context, command *CreateVehicleCommand) {
	vehicleDto := &models.Vehicle{VehicleID: command.VehicleID, Name: command.Name, Description: command.Description, Price: command.Price}
	_, err := createVehicleHandler.pgRepo.CreateVehicle(ctx, vehicleDto)
	if err != nil {
		log.Fatal(err)
	}
}
