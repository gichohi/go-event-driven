package queries

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/repository"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	uuid "github.com/satori/go.uuid"
)

type VehicleQueries struct {
	GetVehicleById GetVehicleByIdHandler
}

func NewVehicleQueries(getVehicleById GetVehicleByIdHandler) *VehicleQueries {
	return &VehicleQueries{GetVehicleById: getVehicleById}
}

type GetVehicleByIdQuery struct {
	VehicleID uuid.UUID `json:"vehicleId" validate:"required,gte=0,lte=255"`
}

func NewGetVehicleByIdQuery(vehicleID uuid.UUID) *GetVehicleByIdQuery {
	return &GetVehicleByIdQuery{VehicleID: vehicleID}
}

type GetVehicleByIdHandler interface {
	Handle(ctx context.Context, query *GetVehicleByIdQuery) (*models.Vehicle, error)
}

type getVehicleByIdHandler struct {
	log    logger.Logger
	cfg    *config.Config
	pgRepo repository.Repository
}

func NewGetVehicleByIdHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository) *getVehicleByIdHandler {
	return &getVehicleByIdHandler{log: log, cfg: cfg, pgRepo: pgRepo}
}

func (q *getVehicleByIdHandler) Handle(ctx context.Context, query *GetVehicleByIdQuery) (*models.Vehicle, error) {
	return q.pgRepo.GetVehicleById(ctx, query.VehicleID)
}


