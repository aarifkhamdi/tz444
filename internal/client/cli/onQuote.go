package cli

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
	"golang.org/x/crypto/argon2"
)

func solveChallenge(challengeID string, difficulty int) (int, error) {
	for nonce := 1; nonce < 1000000; nonce++ {
		// Создаем входные данные
		input := challengeID + fmt.Sprintf("%d", nonce)

		// Вычисляем Argon2
		hash := argon2.IDKey([]byte(input), []byte(challengeID), 1, 64*1024, 1, 32)

		// Проверяем начинается ли с нулевых битов
		if hasLeadingZeros(hash, difficulty) {
			return nonce, nil
		}
	}

	return 0, fmt.Errorf("solution not found within 1000000 attempts")
}

func hasLeadingZeros(hash []byte, difficulty int) bool {
	// Проверяем первые difficulty битов
	for i := 0; i < difficulty/8; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	// Проверяем оставшиеся биты
	if difficulty%8 > 0 {
		mask := byte(0xFF >> (8 - difficulty%8))
		return (hash[difficulty/8] & mask) == 0
	}
	return true
}

func (c *CLI) onQuote() {
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

	nonce, err := solveChallenge(result.ID, result.Difficulty)
	if err != nil {
		zap.L().Error("Failed to solve challenge", zap.Error(err))
		return
	}
	zap.L().Info("Challenge solved", zap.Int("nonce", nonce))

	if c.cfg.SendWrongChallenge {
		nonce++
	}

	err = protocol.WriteRequest(c.conn, &protocol.Request{
		ID:     c.getID(),
		Method: "quote",
		Auth: &protocol.Auth{
			ChallengeID: result.ID,
			Nonce:       nonce,
		},
	})
	if err != nil {
		zap.L().Error("Failed to send quote request", zap.Error(err))
		return
	}
	zap.L().Info("Quote request sent successfully")
	c.handleResponse()
}
