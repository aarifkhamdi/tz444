package handler

import "github.com/aarifkhamdi/tz444/internal/shared/protocol"

func echoHandler(req *protocol.Request) *protocol.Response {
	return &protocol.Response{
		ID:     req.ID,
		Result: req.Params,
	}
}
