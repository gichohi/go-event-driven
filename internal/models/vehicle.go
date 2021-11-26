package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)


type Vehicle struct {
	VehicleID   uuid.UUID `json:"vehicleId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Make 		string	  `json:"make"`
	Model 		string    `json:"model"`
	Year		int64	  `json:"year"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
