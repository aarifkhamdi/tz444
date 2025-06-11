package cli

import (
	"encoding/json"

	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
)

func (c *CLI) onEcho() {
	msg := "Hello, world!"
	var params json.RawMessage
	params, err := json.Marshal(msg)
	if err != nil {
		zap.L().Error("Failed to marshal echo parameters", zap.Error(err))
		return
	}

	zap.L().Info("Sending echo request")
	err = protocol.WriteRequest(c.conn, &protocol.Request{
		ID:     c.getID(),
		Method: "echo",
		Params: params,
	})
	if err != nil {
		zap.L().Error("Failed to send echo request", zap.Error(err))
		return
	}
	zap.L().Info("Echo request sent successfully")
	c.handleResponse()
}
