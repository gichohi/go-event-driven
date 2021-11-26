package dto

import uuid "github.com/satori/go.uuid"

type VehicleDto struct {
	VehicleID   uuid.UUID `json:"VehicleId" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=0,lte=255"`
	Description string    `json:"description" validate:"required,gte=0,lte=5000"`
	Make 		string	  `json:"make" validate:"required,gte=0,lte=255"`
	Model 		string    `json:"model" validate:"required,gte=0,lte=255"`
	Year		int64	  `json:"year" validate:"required,gte=1980,lte=2022"`
	Price       float64   `json:"price" validate:"required,gte=0"`
}
