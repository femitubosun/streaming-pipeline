package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	KafkaBroker string `env:"KAFKA_BROKER" envDefault:"localhost:9092"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
