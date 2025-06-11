package cli

import (
	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
)

func (c *CLI) onError() {
	zap.L().Info("Sending ping request")
	err := protocol.WriteRequest(c.conn, &protocol.Request{
		ID:     c.getID(),
		Method: "error",
	})
	if err != nil {
		zap.L().Error("Failed to send ping request", zap.Error(err))
		return
	}
	zap.L().Info("Ping request sent successfully")
	c.handleResponse()
}
