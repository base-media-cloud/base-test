package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Host              string        `env:"HOST,required"`
	DeliverySvc       string        `env:"DELIVERY_SERVICE,required"`
	Price             float64       `env:"PRICE,required"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT,default=15s"`
}

func New() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
