package ves_websocket

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

type ProcessMessageHandler func(message []byte, messageID wsrpc.MessageType)

type VESWSSocket struct {
	SocketConn
	logger           logger.Logger
	OnProcessMessage ProcessMessageHandler
}

func NewVESSocket(conn SocketConn, handler ProcessMessageHandler, options ...interface{}) (
	VESWSSocket, error) {
	if len(options) != 1 {
		return VESWSSocket{}, errors.New("must with a logger")
	}
	if socLogger, ok := options[0].(logger.Logger); ok {
		return VESWSSocket{SocketConn: conn, OnProcessMessage: handler, logger: socLogger}, nil
	}
	return VESWSSocket{}, errors.New("must with a logger")
}

func Raw(conn SocketConn) VESWSSocket {
	return VESWSSocket{SocketConn: conn}
}

func (soc VESWSSocket) ProcessBinaryMessage(rawMessage []byte) (message []byte, messageID wsrpc.MessageType, err error) {
	message, messageID, err = wsrpc.Deserialize(rawMessage)
	if err != nil {
		err = wrapper.Wrap(types.CodeReadMessageIDError, err)
		return
	}
	return
}

func (soc VESWSSocket) postMessage(buf *bytes.Buffer, err error) error {
	if err != nil {
		return err
	}
	err = soc.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return err
	}
	wsrpc.GetDefaultSerializer().Put(buf)
	return nil
}

func (soc VESWSSocket) PostMessage(code wsrpc.MessageType, msg proto.Message) error {
	return soc.postMessage(wsrpc.GetDefaultSerializer().Serialize(code, msg))
}

func (soc VESWSSocket) PostRawPacket(code wsrpc.MessageType, msg []byte) error {
	return soc.postMessage(wsrpc.GetDefaultSerializer().SerializeRaw(code, msg))
}

func (soc VESWSSocket) ReadRoutine() {
	defer func() {
		fmt.Println("exited read routine")
	}()

	var (
		messageType int
		message     []byte
		err         error
		messageID   wsrpc.MessageType
	)
	for {
		messageType, message, err = soc.ReadMessage()
		if err != nil {
			err = wrapper.Wrap(types.CodeReadMessageError, err)
			return
		}
		switch messageType {
		case websocket.BinaryMessage, websocket.TextMessage:
			message, messageID, err = soc.ProcessBinaryMessage(message)
			if err != nil {
				soc.logger.Error("read message error", err)
				continue
			}

			soc.OnProcessMessage(message, messageID)
		case websocket.PingMessage, websocket.PongMessage:
		case websocket.CloseMessage:
			return
		}
	}
}
