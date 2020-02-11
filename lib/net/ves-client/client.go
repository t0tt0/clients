package vesclient

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/config"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"io"

	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"

	ethbni "github.com/Myriad-Dreamin/go-ves/lib/bni/eth"
	nsbbni "github.com/Myriad-Dreamin/go-ves/lib/bni/ten"
	"github.com/Myriad-Dreamin/go-ves/lib/database/filedb"
)

func (vc *VesClient) load(dbpath string) error {
	var err, err2 error
	filedb.Register(&ECCKeys{})
	filedb.Register(&EthAccounts{})
	if vc.fdb, err = filedb.NewFileDB(dbpath); err != nil {
		return err
	}
	var ev *filedb.ReadEvent
	ev, err = vc.fdb.ReadWithPath("/keys.dat")
	if err != nil {
		goto bad_load_keys
	}
	vc.keys = new(ECCKeys)
	vc.keys.Alias = make(map[string]ECCKey)
	err = ev.Decode(vc.keys)
	if err != nil {
		goto bad_load_keys
	}
	err = ev.Settle()
	if err != nil {
		goto bad_load_keys
	}
bad_load_keys:
	if err == io.EOF {
		err = nil
	}

	ev, err2 = vc.fdb.ReadWithPath("/accs.dat")
	if err2 != nil {
		goto bad_load_accs
	}
	vc.accs = new(EthAccounts)
	vc.accs.Alias = make(map[string]EthAccount)
	err2 = ev.Decode(&vc.accs)
	if err2 != nil {
		goto bad_load_accs
	}
	err2 = ev.Settle()
	if err2 != nil {
		goto bad_load_accs
	}

	return err
bad_load_accs:
	if err2 == io.EOF {
		err2 = err
	} else {
		if err != nil {
			err2 = fmt.Errorf("error loading keys: %v, error loading accs: %v", err, err2)
		}
	}
	return err2
}

func (vc *VesClient) save() {
	var err error
	if err = vc.updateKeys(); err != nil {
		vc.logger.Error("update keys failed", "error", err)
	}
	if err = vc.updateAccs(); err != nil {
		vc.logger.Error("update accounts failed", "error", err)
	}
}

func (vc *VesClient) updateFileObj(name string, obj interface{}) error {
	ev, err := vc.fdb.WriteWithPath(name)
	if err != nil {
		return err
	}
	err = ev.Encode(obj)
	if err != nil {
		err2 := ev.Settle()
		return errors.New(err.Error() + "\n" + err2.Error())
	}
	err = ev.Settle()
	if err != nil {
		return err
	}
	return nil
}

func (vc *VesClient) updateKeys() error {
	return vc.updateFileObj("/keys.dat", vc.keys)
}

func (vc *VesClient) updateAccs() error {
	return vc.updateFileObj("/accs.dat", vc.accs)
}

func (vc *VesClient) getClientHello() *wsrpc.ClientHello {
	return new(wsrpc.ClientHello)
}

func (vc *VesClient) getClientHelloReply() *wsrpc.ClientHelloReply {
	return new(wsrpc.ClientHelloReply)
}

func (vc *VesClient) getShortSendMessage() *wsrpc.Message {
	return new(wsrpc.Message)
}

func (vc *VesClient) getRawMessage() *wsrpc.RawMessage {
	return new(wsrpc.RawMessage)
}

func (vc *VesClient) getShortReplyMessage() *wsrpc.Message {
	return new(wsrpc.Message)
}

func (vc *VesClient) getUserRegisterRequest() *wsrpc.UserRegisterRequest {
	return new(wsrpc.UserRegisterRequest)
}

func (vc *VesClient) getUserRegisterReply() *wsrpc.UserRegisterReply {
	return new(wsrpc.UserRegisterReply)
}

func (vc *VesClient) getrequestComingRequest() *wsrpc.RequestComingRequest {
	return new(wsrpc.RequestComingRequest)
}

func (vc *VesClient) getrequestComingReply() *wsrpc.RequestComingReply {
	return new(wsrpc.RequestComingReply)
}

func (vc *VesClient) getrequestGrpcServiceRequest() *wsrpc.RequestGrpcServiceRequest {
	return new(wsrpc.RequestGrpcServiceRequest)
}

func (vc *VesClient) getrequestGrpcServiceReply() *wsrpc.RequestGrpcServiceReply {
	return new(wsrpc.RequestGrpcServiceReply)
}

func (vc *VesClient) getrequestNsbServiceRequest() *wsrpc.RequestNsbServiceRequest {
	return new(wsrpc.RequestNsbServiceRequest)
}

func (vc *VesClient) getrequestNsbServiceReply() *wsrpc.RequestNsbServiceReply {
	return new(wsrpc.RequestNsbServiceReply)
}

func (vc *VesClient) getuserRegisterRequest() *wsrpc.UserRegisterRequest {
	return new(wsrpc.UserRegisterRequest)
}

func (vc *VesClient) getuserRegisterReply() *wsrpc.UserRegisterReply {
	return new(wsrpc.UserRegisterReply)
}

func (vc *VesClient) getsessionListRequest() *wsrpc.SessionListRequest {
	return new(wsrpc.SessionListRequest)
}

func (vc *VesClient) getsessionListReply() *wsrpc.SessionListReply {
	return new(wsrpc.SessionListReply)
}

func (vc *VesClient) gettransactionListRequest() *wsrpc.TransactionListRequest {
	return new(wsrpc.TransactionListRequest)
}

func (vc *VesClient) gettransactionListReply() *wsrpc.TransactionListReply {
	return new(wsrpc.TransactionListReply)
}

func (vc *VesClient) getSessionStart() *uiprpc.SessionStartRequest {
	return new(uiprpc.SessionStartRequest)
}

func (vc *VesClient) getSessionFinishedRequest() *wsrpc.SessionFinishedRequest {
	return new(wsrpc.SessionFinishedRequest)
}

func (vc *VesClient) getSessionFinishedReply() *wsrpc.SessionFinishedReply {
	return new(wsrpc.SessionFinishedReply)
}

// func (vc *VesClient) getSessionRequireTransactRequest() *wsrpc.SessionRequireTransactRequest {
// 	if vc.sessionRequireTransactRequest == nil {
// 		vc.sessionRequireTransactRequest = new(wsrpc.SessionRequireTransactRequest)
// 	}
// 	return vc.sessionRequireTransactRequest
// }
//
// func (vc *VesClient) getSessionRequireTransactReply() *wsrpc.SessionRequireTransactReply {
// 	if vc.sessionRequireTransactReply == nil {
// 		vc.sessionRequireTransactReply = new(wsrpc.SessionRequireTransactReply)
// 	}
// 	return vc.sessionRequireTransactReply
// }

func (vc *VesClient) getSendAttestationReceiveRequest() *wsrpc.AttestationReceiveRequest {
	// if vc.attestationReceiveRequestSend == nil {
	// 	vc.attestationReceiveRequestSend = new(wsrpc.AttestationReceiveRequest)
	// }
	// return vc.attestationReceiveRequestSend
	return new(wsrpc.AttestationReceiveRequest)
}

func (vc *VesClient) getReceiveAttestationReceiveRequest() *wsrpc.AttestationReceiveRequest {
	// if vc.attestationReceiveRequestReceive == nil {
	// 	vc.attestationReceiveRequestReceive = new(wsrpc.AttestationReceiveRequest)
	// }
	// return vc.attestationReceiveRequestReceive
	return new(wsrpc.AttestationReceiveRequest)
}

func (vc *VesClient) getAttestationReceiveReply() *wsrpc.AttestationReceiveReply {
	return new(wsrpc.AttestationReceiveReply)
}

func (vc *VesClient) getCloseSessionRequest() *wsrpc.CloseSessionRequest {
	return new(wsrpc.CloseSessionRequest)
}

func (vc *VesClient) postMessage(code wsrpc.MessageType, msg proto.Message) error {
	buf, err := wsrpc.GetDefaultSerializer().Serial(code, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	vc.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	wsrpc.GetDefaultSerializer().Put(buf)
	return nil
}

func (vc *VesClient) postRawMessage(code wsrpc.MessageType, dst *uiprpc_base.Account, msg proto.Message) error {

	buf, err := wsrpc.GetDefaultSerializer().Serial(code, msg)
	/// fmt.Println(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}
	var s = vc.getRawMessage()
	s.To, err = proto.Marshal(dst)
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.From = vc.name
	s.Contents = make([]byte, buf.Len())
	copy(s.Contents, buf.Bytes())
	// fmt.Println(s.Contents)
	wsrpc.GetDefaultSerializer().Put(buf)
	return vc.postMessage(wsrpc.CodeRawProto, s)
}

func (vc *VesClient) SendMessage(to, msg []byte) error {
	shortSendMessage := vc.getShortSendMessage()
	shortSendMessage.From = vc.name
	shortSendMessage.To = to
	shortSendMessage.Contents = string(msg)

	// fmt.Println(to, msg)

	return vc.postMessage(wsrpc.CodeMessageRequest, shortSendMessage)
}

func (vc *VesClient) setName(b []byte) {
	vc.rwMutex.Lock()
	defer vc.rwMutex.Unlock()
	vc.name = make([]byte, len(b))
	copy(vc.name, b)
}

func (vc *VesClient) getName() []byte {
	vc.rwMutex.RLock()
	defer vc.rwMutex.RUnlock()
	return vc.name
}

func (vc *VesClient) getNSBSigner() (uiptypes.Signer, error) {
	if vc.signer != nil {
		return vc.signer, nil
	}

	if key, ok := vc.keys.Alias["ten1"]; ok {
		var err error
		vc.signer, err = signaturer.NewTendermintNSBSigner(key.PrivateKey)
		if err != nil {
			return nil, err
		}
		if vc.signer == nil {
			return nil, errIlegalPrivateKey
		}
	} else {
		return nil, errTen1NotFound
	}
	return vc.signer, nil
}

func (vc *VesClient) getRouter(chainID uint64) uiptypes.Router {
	switch chainID {
	case 1, 2:
		return ethbni.NewBN(config.ChainDNS)
	case 3, 4, 5:
		return nsbbni.NewBN(config.ChainDNS)
	default:
		return nil
	}
}

func (vc *VesClient) getBlockStorage(chainID uint64) uiptypes.Storage {
	switch chainID {
	case 1, 2:
		return ethbni.NewBN(config.ChainDNS)
	case 3, 4, 5:
		return nsbbni.NewBN(config.ChainDNS)
	default:
		return nil
	}
}
