package main

import (
	"github.com/aarifkhamdi/tz444/internal/client/cli"
	"github.com/aarifkhamdi/tz444/internal/client/client"
	"github.com/aarifkhamdi/tz444/internal/client/config"
	"github.com/aarifkhamdi/tz444/internal/shared/logger"
	"go.uber.org/zap"
)

func main() {
	defer logger.Init()()

	config := config.New()

	client, err := client.New(config)
	if err != nil {
		zap.L().Fatal("Failed to connect to server", zap.Error(err))
	}
	defer client.Close()

	if err = cli.New(client, config); err != nil {
		zap.L().Error("CLI closed", zap.Error(err))
	}

	zap.L().Info("CLI closed")
}
