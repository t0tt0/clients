package client

import (
	"fmt"
	base_account "github.com/HyperService-Consortium/go-uip/base-account"
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/hub"
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
			c.Hub.Server.Logger.Info("unmarshal error", "error", err)
			return
		}
		c.Hub.Server.Logger.Info("message request",
			"from", "todo", "to", s.GetTo())
		c.Hub.Broadcast <- hub.NewWriteMessageTask(wsrpc.CodeMessageReply, &s)
	case wsrpc.CodeRawProto:

		var s wsrpc.RawMessage
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.Hub.Server.Logger.Info("error", err)
			return
		}
		c.Hub.Server.Logger.Info("raw proto",
			"from", "todo", "to", s.GetTo())

		c.Hub.Unicast <- &hub.UniMessage{Target: s.GetTo(), Task: hub.NewRawWriteMessageTask(
			wsrpc.MessageType(s.MessageType),
			s.GetContents())}
	case wsrpc.CodeClientHelloRequest:
		var s wsrpc.ClientHello
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.Hub.Server.Logger.Info("error", err)
			return
		}

		c.User, err = c.Hub.Server.UserDB.FindUser(string(s.GetName()))
		// fmt.Println(c.user, err)
		if err != nil {
			c.Hub.Server.Logger.Error("find user error", "error", err)
			return
		} else if c.User == nil {
			c.Hub.Server.Logger.Error("user not found", "error", err)
			return
		}

		select {
		case c.Helloed <- true:
			var t wsrpc.ClientHelloReply
			t.GrpcHost = gRpcIPs[0]
			t.NsbHost = c.Hub.Server.Nsbip
			c.Hub.Unicast <- &hub.UniMessage{Target: &base_account.Account{
				ChainId: hub.PlaceHolderChain, Address: s.GetName(),
			}, Task: hub.NewWriteMessageTask(wsrpc.CodeClientHelloReply, &t)}
		default:
		}

	case wsrpc.CodeUserRegisterRequest:
		var s wsrpc.UserRegisterRequest
		err = proto.Unmarshal(message, &s)
		if err != nil {
			c.Hub.Server.Logger.Info("error", err)
		}

		// fmt.Println("hexx registering", hex.EncodeToString(s.GetAccount().GetAddress()))
		err = c.Hub.Server.UserDB.InsertAccount(s.GetUserName(), s.GetAccount())

		if err != nil {
			c.Hub.Server.Logger.Info("error", err)
			return
		}
	default:
		fmt.Println("aborting message", string(message))
		// abort
	}

	// c.hub.broadcast <- &broMessage{bytes.TrimSpace(bytes.Replace(message, newline, space, -1)), func() {}}
}

