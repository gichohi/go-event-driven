package handlers

import (
	"encoding/json"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/dto"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/events"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/models"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func NewHandler() http.Handler {
	router := chi.NewRouter()
	router.Route("/vehicles", routes)
	return router
}

func routes(router chi.Router) {
	router.Post("/", CreateVehicle)
	router.Get("/", Top)

}

func Top(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("<h1>Welcome to my server</h1>"))
	if err != nil {
		log.Fatal(err)
	}

}

func CreateVehicle(w http.ResponseWriter, r *http.Request){
	var vehicle models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&vehicle)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createDto := &dto.VehicleDto{}
	createDto.VehicleID = vehicle.VehicleID
	createDto.Name = vehicle.Name
	createDto.Model = vehicle.Model
	createDto.Make = vehicle.Make
	createDto.Description = vehicle.Description
	createDto.Year = vehicle.Year
	createDto.Price = vehicle.Price

	var v = events.VehicleCreatedHandler{}
	go func() {
		v.Handle(events.NewVehicleCreatedEvent(createDto))
	}()

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("<h1>Created</h1>"))
	if err != nil {
		log.Fatalf("Write Error: %s", err)
	}
}
