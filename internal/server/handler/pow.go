package handler

import (
	"github.com/aarifkhamdi/tz444/internal/server/challenge"
	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
)

type challengeGenerator interface {
	GenerateChallenge() (*challenge.Challenge, error)
}

type challengeVerifier interface {
	VerifySolution(challengeID string, nonce int) (bool, error)
}

func generateChallengeHandler(
	challengeGenerator challengeGenerator,
) handler {
	return func(req *protocol.Request) *protocol.Response {
		challenge, err := challengeGenerator.GenerateChallenge()
		if err != nil {
			zap.L().Error("Failed to generate challenge", zap.Error(err))
			return protocol.NewErrorResponse(req.ID, 500, "Failed to generate challenge")
		}

		resp, err := protocol.NewResponse(req.ID, challenge)
		if err != nil {
			zap.L().Error("Failed to create response", zap.Error(err))
			return protocol.NewErrorResponse(req.ID, 500, "Failed to create response")
		}

		return resp
	}
}

func verifyChallengeMiddleware(
	next handler,
	challengeVerifier challengeVerifier,
) handler {
	return func(req *protocol.Request) *protocol.Response {
		if req.Auth == nil {
			return protocol.NewErrorResponse(req.ID, 401, "Unauthorized")
		}
		ok, err := challengeVerifier.VerifySolution(req.Auth.ChallengeID, req.Auth.Nonce)
		if err != nil {
			zap.L().Error("Failed to verify solution", zap.Error(err))
			return protocol.NewErrorResponse(req.ID, 500, "Failed to verify solution")
		}
		if !ok {
			return protocol.NewErrorResponse(req.ID, 401, "Unauthorized")
		}

		return next(req)
	}
}
