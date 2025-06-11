package handler

import "github.com/aarifkhamdi/tz444/internal/shared/protocol"

func errorHandler(req *protocol.Request) *protocol.Response {
	return protocol.NewErrorResponse(req.ID, 400, "example error response")
}
