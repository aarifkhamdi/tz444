package handler

import (
	"github.com/aarifkhamdi/tz444/internal/server/quote"
	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
)

func quoteHandler(req *protocol.Request) *protocol.Response {
	quote := quote.GetRandomQuote()
	resp, err := protocol.NewResponse(req.ID, quote)
	if err != nil {
		return protocol.NewErrorResponse(req.ID, 500, "Failed to create response")
	}

	return resp
}
