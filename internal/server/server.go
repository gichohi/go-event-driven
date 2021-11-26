package server

import (
	"github.com/gichohi/go-cqrs-kafka-grpc/config"
	"github.com/gichohi/go-cqrs-kafka-grpc/internal/service"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/postgres"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)


type server struct {
	log       logger.Logger
	cfg       *config.Config
	v         *validator.Validate
	kafkaConn *kafka.Conn
	ps        *service.VehicleService
	pgConn    *pgxpool.Pool
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	pgxConn, err := postgres.NewPgxConn(s.cfg.Postgresql)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.pgConn = pgxConn
	s.log.Infof("postgres connected: %v", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()

	return nil
}



