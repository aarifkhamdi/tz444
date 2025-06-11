package config

import (
	"github.com/aarifkhamdi/tz444/internal/shared/config"
)

type Config struct {
	Addr               string `mapstructure:"addr" env:"ADDR" validate:"required"`
	IsInteractive      bool   `mapstructure:"isInteractive" env:"INTERACTIVE"`
	SendWrongChallenge bool   `mapstructure:"sendWrongChallenge" env:"SEND_WRONG_CHALLENGE"`
}

func New() *Config {
	cfg := &Config{}
	return config.New(cfg)
}
