package handler

import (
	"github.com/aarifkhamdi/tz444/internal/server/challenge"
	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
)

type handler func(*protocol.Request) *protocol.Response

type RequestHandler struct {
	handlersMap map[string]handler
}

func New() *RequestHandler {
	powService := challenge.NewPoWService()

	handlersMap := map[string]handler{
		"challenge": generateChallengeHandler(powService),
		"quote":     verifyChallengeMiddleware(quoteHandler, powService),
		"echo":      echoHandler,
		"error":     errorHandler,
	}

	reqHandler := &RequestHandler{
		handlersMap: handlersMap,
	}

	return reqHandler
}

func (h *RequestHandler) HandleRequest(req *protocol.Request) *protocol.Response {
	handler := h.handlersMap[req.Method]
	if handler == nil {
		return protocol.NewErrorResponse(req.ID, 404, "Method not found")
	}
	return handler(req)
}
