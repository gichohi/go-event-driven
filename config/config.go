package config

import (
	"flag"
	"fmt"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/constants"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/logger"
	"github.com/gichohi/go-cqrs-kafka-grpc/pkg/postgres"
	"github.com/pkg/errors"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Writer microservice microservice config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	Logger      *logger.Config      `mapstructure:"logger"`
	GRPC        GRPC                `mapstructure:"grpc"`
	Postgresql  *postgres.Config    `mapstructure:"postgres"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/writer_service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}


	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}

	kafkaUrl := os.Getenv(constants.KafkaUrl)
	if kafkaUrl != "" {
		cfg.GRPC.Port = grpcPort
	}

	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}

	return cfg, nil
}
