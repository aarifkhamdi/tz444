package main

import (
	"errors"
	"net"

	"github.com/aarifkhamdi/tz444/internal/server/config"
	"github.com/aarifkhamdi/tz444/internal/server/server"
	"github.com/aarifkhamdi/tz444/internal/shared/logger"
	"go.uber.org/zap"
)

func main() {
	defer logger.Init()()

	cfg := config.New()

	srv := server.NewServer(cfg)

	zap.L().Info("Starting TCP server", zap.String("address", cfg.Addr))

	if err := srv.Start(); err != nil {
		if errors.Is(err, net.ErrClosed) {
			zap.L().Info("Server closed", zap.Error(err))
			return
		}

		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}
