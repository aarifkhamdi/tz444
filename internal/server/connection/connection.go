package connection

import (
	"net"
	"sync"

	"github.com/aarifkhamdi/tz444/internal/server/handler"
	"github.com/aarifkhamdi/tz444/internal/shared/protocol"
	"go.uber.org/zap"
)

type ConnectionHandler struct {
	conn           net.Conn
	requestHandler *handler.RequestHandler
	writeMu        sync.Mutex
}

func New(conn net.Conn, requestHandler *handler.RequestHandler) *ConnectionHandler {
	connectionHandler := &ConnectionHandler{
		conn:           conn,
		requestHandler: requestHandler,
	}

	go connectionHandler.run()

	return connectionHandler
}

func (ch *ConnectionHandler) run() {
	clientAddr := ch.conn.RemoteAddr().String()

	for {
		req, err := protocol.ReadRequest(ch.conn)
		if err != nil {
			zap.L().Info("Client disconnected", zap.String("address", clientAddr), zap.Error(err))
			return
		}

		zap.L().Info(
			"Received request",
			zap.String("address", clientAddr),
			zap.Int64("id", req.ID),
			zap.String("method", req.Method),
		)

		go ch.handleRequestAsync(req, clientAddr)
	}
}

func (ch *ConnectionHandler) handleRequestAsync(req *protocol.Request, clientAddr string) {
	response := ch.requestHandler.HandleRequest(req)

	ch.writeMu.Lock()
	err := protocol.WriteResponse(ch.conn, response)
	ch.writeMu.Unlock()

	if err == nil {
		zap.L().Info(
			"Sent response",
			zap.String("address", clientAddr),
			zap.Int64("id", req.ID),
		)
	} else {
		zap.L().Error(
			"Failed to send response",
			zap.String("address", clientAddr),
			zap.Int64("id", req.ID),
			zap.Error(err),
		)
	}
}
