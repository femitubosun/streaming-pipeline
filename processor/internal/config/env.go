package config

import (
	"os"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	KafkaBrokers string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	ProcessorID  string `env:"PROCESSOR_ID"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	if cfg.ProcessorID == "" {
		hostname, _ := os.Hostname()
		cfg.ProcessorID = hostname
	}

	return &cfg, nil
}
