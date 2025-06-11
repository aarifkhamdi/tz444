package server

import (
	"errors"
	"net"

	"github.com/aarifkhamdi/tz444/internal/server/config"
	"github.com/aarifkhamdi/tz444/internal/server/connection"
	"github.com/aarifkhamdi/tz444/internal/server/handler"
	"go.uber.org/zap"
)

type Server struct {
	address string
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		address: cfg.Addr,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	zap.L().Info("Server started", zap.String("address", s.address))

	requestHandler := handler.New()

	for {
		conn, err := listener.Accept()

		if errors.Is(err, net.ErrClosed) {
			return err
		}

		if err != nil {
			zap.L().Error("Failed to accept connection", zap.Error(err))
			continue
		}

		clientAddr := conn.RemoteAddr().String()
		zap.L().Info("New client connected", zap.String("address", clientAddr))

		connection.New(conn, requestHandler)
	}
}
