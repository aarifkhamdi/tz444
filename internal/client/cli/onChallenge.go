package cli

import (
	"encoding/json"
	"time"

	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
)

func (c *CLI) onChallenge() {
	zap.L().Info("Sending pow request")
	err := protocol.WriteRequest(c.conn, &protocol.Request{
		ID:     c.getID(),
		Method: "challenge",
	})
	if err != nil {
		zap.L().Error("Failed to send pow request", zap.Error(err))
		return
	}
	zap.L().Info("Pow request sent successfully")
	resp, err := c.handleResponse()
	if err != nil {
		zap.L().Error("Failed to handle pow response", zap.Error(err))
		return
	}
	type challenge struct {
		ID         string    `json:"id"`
		Difficulty int       `json:"difficulty"`
		ExpiresAt  time.Time `json:"expires_at"`
	}
	var result challenge
	if err := json.Unmarshal(resp.Result, &result); err != nil {
		zap.L().Error("Failed to unmarshal pow result", zap.Error(err))
		return
	}
	zap.L().Info(
		"Pow result received",
		zap.String("challenge", result.ID),
		zap.Int("difficulty", result.Difficulty),
	)
}
