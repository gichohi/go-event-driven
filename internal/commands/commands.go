package commands

import (
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/dto"
	uuid "github.com/satori/go.uuid"
)

type VehicleCommands struct {
	CreateVehicle CreateVehicleCmdHandler
	UpdateVehicle UpdateVehicleCmdHandler
	DeleteVehicle DeleteVehicleCmdHandler
}

func NewVehicleCommands(createVehicle CreateVehicleCmdHandler, updateVehicle UpdateVehicleCmdHandler, deleteVehicle DeleteVehicleCmdHandler) *VehicleCommands {
	return &VehicleCommands{CreateVehicle: createVehicle, UpdateVehicle: updateVehicle, DeleteVehicle: deleteVehicle}
}

type CreateVehicleCommand struct {
	VehicleID   uuid.UUID `json:"VehicleId" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=0,lte=255"`
	Description string    `json:"description" validate:"required,gte=0,lte=5000"`
	Make 		string	  `json:"make" validate:"required,gte=0,lte=255"`
	Model 		string    `json:"model" validate:"required,gte=0,lte=255"`
	Year		int64	  `json:"year" validate:"required,gte=1980,lte=2022"`
	Price       float64   `json:"price" validate:"required,gte=0"`
}

func NewCreateVehicleCommand(v *dto.VehicleDto) *CreateVehicleCommand {
	return &CreateVehicleCommand{VehicleID: v.VehicleID, Name: v.Name, Description: v.Description, Make: v.Make, Model: v.Model, Year: v.Year, Price: v.Price}
}

type UpdateVehicleCommand struct {
	VehicleID   uuid.UUID `json:"VehicleId" validate:"required,gte=0,lte=255"`
	Name        string    `json:"name" validate:"required,gte=0,lte=255"`
	Description string    `json:"description" validate:"required,gte=0,lte=5000"`
	Make 		string	  `json:"make" validate:"required,gte=0,lte=255"`
	Model 		string    `json:"model" validate:"required,gte=0,lte=255"`
	Year		int64	  `json:"year" validate:"required,gte=1980,lte=2022"`
	Price       float64   `json:"price" validate:"required,gte=0"`
}

func NewUpdateVehicleCommand(productID uuid.UUID, name string, description string, make string, model string, year int64, price float64) *UpdateVehicleCommand {
	return &UpdateVehicleCommand{VehicleID: productID, Name: name, Description: description, Make: make, Model: model, Year: year, Price: price}
}

type DeleteVehicleCommand struct {
	VehicleID uuid.UUID `json:"VehicleId" validate:"required"`
}

func NewDeleteVehicleCommand(productID uuid.UUID) *DeleteVehicleCommand {
	return &DeleteVehicleCommand{VehicleID: productID}
}

type ReadVehicleCommand struct {
	VehicleID uuid.UUID `json:"VehicleId" validate:"required"`
}
