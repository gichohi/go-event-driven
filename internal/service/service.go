package service

import (
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/commands"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/events"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/queries"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/repository"
	KafkaClient "github.com/gichohi/go-cqrs-kafka-grpc/pkg/kafka"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
)


type params struct {
	log  logger.Logger
	cfg  *config.Config
	pgRepo  repository.Repository
	producer KafkaClient.Producer
}

type VehicleService struct {
	Commands *commands.VehicleCommands
	Events 	 *events.VehicleEvents
	Queries  *queries.VehicleQueries
}

func NewVehicleService() *VehicleService {
	var p *params

	VehicleCreateCommandHandler := commands.NewCreateVehicleHandler(p.log,p.cfg)
	VehicleUpdateCommandHandler := commands.NewUpdateVehicleHandler(p.log,p.cfg)
	VehicleDeleteCommandHandler := commands.NewDeleteVehicleHandler(p.log,p.cfg)

	VehicleCreateEventHandler := events.NewVehicleCreatedHandler(p.log, p.cfg, p.producer)
	VehicleUpdatedEventHandler := events.NewVehicleUpdatedHandler(p.log, p.cfg)
	VehicleDeletedEventHandler := events.NewDeletedVehicleHandler(p.log, p.cfg)

	getVehicleByIdHandler := queries.NewGetVehicleByIdHandler(p.log, p.cfg, p.pgRepo)

	vehicleCommands := commands.NewVehicleCommands(VehicleCreateCommandHandler, VehicleUpdateCommandHandler, VehicleDeleteCommandHandler)
	vehicleEvents := events.NewVehicleEvents(VehicleCreateEventHandler, VehicleUpdatedEventHandler, VehicleDeletedEventHandler)
	vehicleQueries := queries.NewVehicleQueries(getVehicleByIdHandler)

	return &VehicleService{Commands: vehicleCommands, Events: vehicleEvents, Queries: vehicleQueries}
}

