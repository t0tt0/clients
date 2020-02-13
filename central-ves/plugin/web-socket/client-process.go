package centered_ves

import (
	"fmt"
	base_account "github.com/HyperService-Consortium/go-uip/base-account"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/gogo/protobuf/proto"
)

func (c *Client) ProcessMessage(message []byte, messageID wsrpc.MessageType) {
	var err error
	switch messageID {
	case wsrpc.CodeMessageRequest:
		var s wsrpc.Message
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.hub.server.logger.Info("unmarshal error", "error", err)
			return
		}
		c.hub.server.logger.Info("message request",
			"from", "todo", "to", s.GetTo())
		c.hub.broadcast <- newWriteMessageTask(wsrpc.CodeMessageReply, &s)
	case wsrpc.CodeRawProto:

		var s wsrpc.RawMessage
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.hub.server.logger.Info("error", err)
			return
		}
		c.hub.server.logger.Info("raw proto",
			"from", "todo", "to", s.GetTo())

		c.hub.unicast <- &uniMessage{target: s.GetTo(), task: newRawWriteMessageTask(
			wsrpc.MessageType(s.MessageType),
			s.GetContents())}
	case wsrpc.CodeClientHelloRequest:
		var s wsrpc.ClientHello
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.hub.server.logger.Info("error", err)
			return
		}

		c.user, err = c.hub.server.userDB.FindUser(string(s.GetName()))
		// fmt.Println(c.user, err)
		if err != nil {
			c.hub.server.logger.Error("find user error", "error", err)
			return
		} else if c.user == nil {
			c.hub.server.logger.Error("user not found", "error", err)
			return
		}

		select {
		case c.helloed <- true:
			var t wsrpc.ClientHelloReply
			t.GrpcHost = gRpcIPs[0]
			t.NsbHost = c.hub.server.nsbip
			c.hub.unicast <- &uniMessage{target: &base_account.Account{
				ChainId: placeHolderChain, Address: s.GetName(),
			}, task: newWriteMessageTask(wsrpc.CodeClientHelloReply, &t)}
		default:
		}

	case wsrpc.CodeUserRegisterRequest:
		var s wsrpc.UserRegisterRequest
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.hub.server.logger.Info("error", err)
		}

		// fmt.Println("hexx registering", hex.EncodeToString(s.GetAccount().GetAddress()))
		err = c.hub.server.userDB.InsertAccount(s.GetUserName(), s.GetAccount())

		if err != nil {
			c.hub.server.logger.Info("error", err)
			return
		}
	default:
		fmt.Println("aborting message", string(message))
		// abort
	}

	// c.hub.broadcast <- &broMessage{bytes.TrimSpace(bytes.Replace(message, newline, space, -1)), func() {}}
}

