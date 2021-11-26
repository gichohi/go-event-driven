package main

import (
	"context"
	"github.com/gichohi/go-cqrs-kafka-grpc/api/handlers"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/kafka"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/postgres"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	kafka.CreateVehicleReader("VehicleCreated")

	pgxConn, err := postgres.NewPgxConn()
	if err != nil {
		log.Printf("postgresql.NewPgxConn")
	}
	log.Printf("postgres connected: %v\n", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()

	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	httpHandler := handlers.NewHandler()
	s := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		err := s.Serve(listener)
		if err != nil {
			log.Fatalf("Error occurred: %s", err.Error())
		} else {
			log.Printf("Listening on port: %s", addr)
		}
	}()

	<-ctx.Done()

	defer Stop(s)
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}