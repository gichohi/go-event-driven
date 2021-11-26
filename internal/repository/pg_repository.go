package repository

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type vehicleRepository struct {
	log logger.Logger
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewVehicleRepository(log logger.Logger, cfg *config.Config, db *pgxpool.Pool) *vehicleRepository {
	return &vehicleRepository{log: log, cfg: cfg, db: db}
}

func (repo *vehicleRepository) CreateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {

	var created models.Vehicle
	if err := repo.db.QueryRow(ctx, createVehicleQuery, &vehicle.VehicleID, &vehicle.Name, &vehicle.Description, &vehicle.Price).Scan(
		&created.VehicleID,
		&created.Name,
		&created.Description,
		&created.Price,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "db.QueryRow")
	}

	return &created, nil
}

func (p *vehicleRepository) UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {

	var prod models.Vehicle
	if err := p.db.QueryRow(
		ctx,
		updateVehicleQuery,
		&vehicle.Name,
		&vehicle.Description,
		&vehicle.Price,
		&vehicle.VehicleID,
	).Scan(&prod.VehicleID, &prod.Name, &prod.Description, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &prod, nil
}

func (p *vehicleRepository) GetVehicleById(ctx context.Context, uuid uuid.UUID) (*models.Vehicle, error) {

	var vehicle models.Vehicle
	if err := p.db.QueryRow(ctx, getVehicleByIdQuery, uuid).Scan(
		&vehicle.VehicleID,
		&vehicle.Name,
		&vehicle.Description,
		&vehicle.Price,
		&vehicle.CreatedAt,
		&vehicle.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &vehicle, nil
}

func (p *vehicleRepository) DeleteVehicleByID(ctx context.Context, uuid uuid.UUID) error {

	_, err := p.db.Exec(ctx, deleteVehicleByIdQuery, uuid)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}

	return nil
}
