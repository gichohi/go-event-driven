package events

import (
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/dto"
	uuid "github.com/satori/go.uuid"
)

type VehicleEvents struct {
	CreateVehicle VehicleCreatedEventHandler
	UpdateVehicle VehicleUpdatedEventHandler
	DeleteVehicle VehicleDeletedEventHandler
}

func NewVehicleEvents(createVehicle VehicleCreatedEventHandler, updateVehicle VehicleUpdatedEventHandler, deleteVehicle VehicleDeletedEventHandler) *VehicleEvents {
	return &VehicleEvents{CreateVehicle: createVehicle, UpdateVehicle: updateVehicle, DeleteVehicle: deleteVehicle}
}

type VehicleCreatedEvent struct {
	CreateDto *dto.VehicleDto
}

func NewVehicleCreatedEvent(createDto *dto.VehicleDto) *VehicleCreatedEvent {
	return &VehicleCreatedEvent{CreateDto: createDto}
}

type UpdateVehicleEvent struct {
	VehicleID   uuid.UUID `json:"VehicleId" validate:"required,gte=0,lte=255"`
	Name        string    `json:"name" validate:"required,gte=0,lte=255"`
	Description string    `json:"description" validate:"required,gte=0,lte=5000"`
	Make 		string	  `json:"make" validate:"required,gte=0,lte=255"`
	Model 		string    `json:"model" validate:"required,gte=0,lte=255"`
	Year		int64	  `json:"year" validate:"required,gte=1980,lte=2022"`
	Price       float64   `json:"price" validate:"required,gte=0"`
}

func NewUpdateVehicleEvent(productID uuid.UUID, name string, description string, price float64) *UpdateVehicleEvent {
	return &UpdateVehicleEvent{VehicleID: productID, Name: name, Description: description, Price: price}
}

type DeleteVehicleEvent struct {
	VehicleID uuid.UUID `json:"VehicleId" validate:"required"`
}

func NewDeleteVehicleEvent(productID uuid.UUID) *DeleteVehicleEvent {
	return &DeleteVehicleEvent{VehicleID: productID}
}

type ReadVehicleEvent struct {
	VehicleID uuid.UUID `json:"VehicleId" validate:"required"`
}
