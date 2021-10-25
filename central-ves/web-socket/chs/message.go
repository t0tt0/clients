package chs

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
	"github.com/gogo/protobuf/proto"
)

type WriteMessageTask struct {
	MessageType wsrpc.MessageType
	Message     interface{}
}

type BroMessage = WriteMessageTask
type UniMessage struct {
	Target uip.Account
	Task   *WriteMessageTask
}

func NewRawWriteMessageTask(messageID wsrpc.MessageType, msg []byte) *WriteMessageTask {
	return &WriteMessageTask{MessageType: messageID, Message: msg}
}

func NewWriteMessageTask(messageID wsrpc.MessageType, msg proto.Message) *WriteMessageTask {
	return &WriteMessageTask{MessageType: messageID, Message: msg}
}

func NewRawBroMessage(messageID wsrpc.MessageType, msg []byte) *WriteMessageTask {
	return &WriteMessageTask{MessageType: messageID, Message: msg}
}

func NewBroMessage(messageID wsrpc.MessageType, msg proto.Message) *WriteMessageTask {
	return &WriteMessageTask{MessageType: messageID, Message: msg}
}
