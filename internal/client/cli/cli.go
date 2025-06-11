package cli

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/aarifkhamdi/tz444/internal/client/config"
	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
)

type CLI struct {
	conn net.Conn
	id   int64
	cfg  *config.Config
}

func New(conn net.Conn, cfg *config.Config) error {
	cli := &CLI{
		conn: conn,
		cfg:  cfg,
	}

	if cfg.IsInteractive {
		return cli.run()
	}

	cli.onQuote()

	return nil
}

func (h *CLI) getID() int64 {
	h.id++
	return h.id
}

func (h *CLI) handleResponse() (*protocol.Response, error) {
	resp, err := protocol.ReadResponse(h.conn)
	if err != nil {
		zap.L().Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.Error != nil {
		zap.L().Warn(
			"Received error response",
			zap.Int64("id", resp.ID),
			zap.Int("code", resp.Error.Code),
			zap.String("message", resp.Error.Message),
		)
	} else {
		zap.L().Info(
			"Received success response",
			zap.Int64("id", resp.ID),
			zap.Any("result", resp.Result),
		)
	}

	return resp, nil
}

func (c *CLI) run() error {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		command, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		command = strings.TrimSpace(command)

		switch command {
		case "error":
			c.onError()
		case "echo":
			c.onEcho()
		case "challenge":
			c.onChallenge()
		case "quote":
			c.onQuote()
		case "exit":
			return nil
		}
	}
}
