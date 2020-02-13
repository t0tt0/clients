package centered_ves

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/gogo/protobuf/proto"
)

type writeMessageTask struct {
	messageType wsrpc.MessageType
	message     interface{}
}

type broMessage = writeMessageTask
type uniMessage struct {
	target uiptypes.Account
	task   *writeMessageTask
}

func newRawWriteMessageTask(messageID wsrpc.MessageType, msg []byte) *writeMessageTask {
	return &writeMessageTask{messageType:messageID, message:msg}
}

func newWriteMessageTask(messageID wsrpc.MessageType, msg proto.Message) *writeMessageTask {
	return &writeMessageTask{messageType:messageID, message:msg}
}

func newRawBroMessage(messageID wsrpc.MessageType, msg []byte) *writeMessageTask {
	return &writeMessageTask{messageType:messageID, message:msg}
}

func newBroMessage(messageID wsrpc.MessageType, msg proto.Message) *writeMessageTask {
	return &writeMessageTask{messageType:messageID, message:msg}
}

