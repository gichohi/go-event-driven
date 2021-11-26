package repository

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	CreateVehicle(ctx context.Context, product *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(ctx context.Context, product *models.Vehicle) (*models.Vehicle, error)
	DeleteVehicleByID(ctx context.Context, uuid uuid.UUID) error

	GetVehicleById(ctx context.Context, uuid uuid.UUID) (*models.Vehicle, error)
}
