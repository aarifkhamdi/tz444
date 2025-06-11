package config

import (
	"github.com/aarifkhamdi/tz444/internal/shared/config"
)

type Config struct {
	Addr string `mapstructure:"addr" env:"ADDR" validate:"required"`
}

func New() *Config {
	cfg := &Config{}
	return config.New(cfg)
}
