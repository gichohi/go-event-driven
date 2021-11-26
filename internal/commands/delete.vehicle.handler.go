package commands

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/repository"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"log"
)

type DeleteVehicleCmdHandler interface {
	Handle(ctx context.Context, command *DeleteVehicleCommand)
}

type DeleteVehicleHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
}

func NewDeleteVehicleHandler(log logger.Logger, cfg *config.Config) *DeleteVehicleHandler {
	return &DeleteVehicleHandler{log: log, cfg: cfg}
}

func (deleteVehicleHandler *DeleteVehicleHandler) Handle(ctx context.Context, command *DeleteVehicleCommand) {
	err := deleteVehicleHandler.pgRepo.DeleteVehicleByID(ctx, command.VehicleID)
	if err != nil {
		log.Fatal(err)
	}
}
