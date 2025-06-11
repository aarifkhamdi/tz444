package client

import (
	"net"

	"github.com/aarifkhamdi/tz444/internal/client/config"
	"go.uber.org/zap"
)

type Client struct {
	net.Conn
}

func New(cfg *config.Config) (*Client, error) {
	conn, err := net.Dial("tcp", cfg.Addr)
	if err != nil {
		return nil, err
	}

	zap.L().Info("Connected to server", zap.String("address", cfg.Addr))

	client := &Client{
		conn,
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
