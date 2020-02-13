package vesclient

import (
	"context"
	"encoding/hex"
	signaturetype "github.com/HyperService-Consortium/go-uip/const/signature_type"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	uiptypes "github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"
)

// ECCKey is the private key object in memory
type ECCKey struct {
	PrivateKey []byte                         `json:"private_key"`
	ChainID    uiptypes.ChainIDUnderlyingType `json:"chain_id"`
}

// ECCKeyAlias is the private key object in json
type ECCKeyAlias struct {
	PrivateKey string                         `json:"private_key"`
	ChainID    uiptypes.ChainIDUnderlyingType `json:"chain_id"`
	Alias      string                         `json:"alias"`
}

// EthAccount is the account object in memory
type EthAccount struct {
	Address    string                         `json:"address"`
	ChainID    uiptypes.ChainIDUnderlyingType `json:"chain_id"`
	PassPhrase string                         `json:"pass_phrase"`
}

// EthAccountAlias is the account object in json
type EthAccountAlias struct {
	EthAccount
	Alias string `json:"alias"`
}

// ECCKeys is the object saved in files
type ECCKeys struct {
	Keys  []*ECCKey
	Alias map[string]ECCKey
}

// EthAccounts is the object saved in files
type EthAccounts struct {
	Accs  []*EthAccount
	Alias map[string]EthAccount
}

type SocketConn interface {
	// Subprotocol returns the negotiated protocol for the connection.
	Subprotocol() string
	// Close closes the underlying network connection without sending or waiting
	// for a close message.
	Close() error
	// LocalAddr returns the local network address.
	LocalAddr() net.Addr
	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr
	// WriteControl writes a control message with the given deadline. The allowed
	// message types are CloseMessage, PingMessage and PongMessage.
	WriteControl(messageType int, data []byte, deadline time.Time) error
	// NextWriter returns a writer for the next message to send. The writer's Close
	// method flushes the complete message to the network.
	//
	// There can be at most one open writer on a connection. NextWriter closes the
	// previous writer if the application has not already done so.
	//
	// All message types (TextMessage, BinaryMessage, CloseMessage, PingMessage and
	// PongMessage) are supported.
	NextWriter(messageType int) (io.WriteCloser, error)
	// WritePreparedMessage writes prepared message into connection.
	WritePreparedMessage(pm *websocket.PreparedMessage) error
	// WriteMessage is a helper method for getting a writer using NextWriter,
	// writing the message and closing the writer.
	WriteMessage(messageType int, data []byte) error
	// SetWriteDeadline sets the write deadline on the underlying network
	// connection. After a write has timed out, the websocket state is corrupt and
	// all future writes will return an error. A zero value for t means writes will
	// not time out.
	SetWriteDeadline(t time.Time) error
	// NextReader returns the next data message received from the peer. The
	// returned messageType is either TextMessage or BinaryMessage.
	//
	// There can be at most one open reader on a connection. NextReader discards
	// the previous message if the application has not already consumed it.
	//
	// Applications must break out of the application's read loop when this method
	// returns a non-nil error value. Errors returned from this method are
	// permanent. Once this method returns a non-nil error, all subsequent calls to
	// this method return the same error.
	NextReader() (messageType int, r io.Reader, err error)
	// ReadMessage is a helper method for getting a reader using NextReader and
	// reading from that reader to a buffer.
	ReadMessage() (messageType int, p []byte, err error)
	// SetReadDeadline sets the read deadline on the underlying network connection.
	// After a read has timed out, the websocket connection state is corrupt and
	// all future reads will return an error. A zero value for t means reads will
	// not time out.
	SetReadDeadline(t time.Time) error
	// SetReadLimit sets the maximum size for a message read from the peer. If a
	// message exceeds the limit, the connection sends a close message to the peer
	// and returns ErrReadLimit to the application.
	SetReadLimit(limit int64)
	// CloseHandler returns the current close handler
	CloseHandler() func(code int, text string) error
	// SetCloseHandler sets the handler for close messages received from the peer.
	// The code argument to h is the received close code or CloseNoStatusReceived
	// if the close message is empty. The default close handler sends a close
	// message back to the peer.
	//
	// The handler function is called from the NextReader, ReadMessage and message
	// reader Read methods. The application must read the connection to process
	// close messages as described in the section on Control Messages above.
	//
	// The connection read methods return a CloseError when a close message is
	// received. Most applications should handle close messages as part of their
	// normal error handling. Applications should only set a close handler when the
	// application must perform some action before sending a close message back to
	// the peer.
	SetCloseHandler(h func(code int, text string) error)
	// PingHandler returns the current ping handler
	PingHandler() func(appData string) error
	// SetPingHandler sets the handler for ping messages received from the peer.
	// The appData argument to h is the PING message application data. The default
	// ping handler sends a pong to the peer.
	//
	// The handler function is called from the NextReader, ReadMessage and message
	// reader Read methods. The application must read the connection to process
	// ping messages as described in the section on Control Messages above.
	SetPingHandler(h func(appData string) error)
	// PongHandler returns the current pong handler
	PongHandler() func(appData string) error
	// SetPongHandler sets the handler for pong messages received from the peer.
	// The appData argument to h is the PONG message application data. The default
	// pong handler does nothing.
	//
	// The handler function is called from the NextReader, ReadMessage and message
	// reader Read methods. The application must read the connection to process
	// pong messages as described in the section on Control Messages above.
	SetPongHandler(h func(appData string) error)
	// UnderlyingConn returns the internal net.Conn. This can be used to further
	// modifications to connection specific flags.
	UnderlyingConn() net.Conn
	// EnableWriteCompression enables and disables write compression of
	// subsequent text and binary messages. This function is a noop if
	// compression was not negotiated with the peer.
	EnableWriteCompression(enable bool)
	// SetCompressionLevel sets the flate compression level for subsequent text and
	// binary messages. This function is a noop if compression was not negotiated
	// with the peer. See the compress/flate package for a description of
	// compression levels.
	SetCompressionLevel(level int) error
}

func NewEthAccount(publicAddress, passPhrase []byte,
	chainID uiptypes.ChainIDUnderlyingType) (EthAccount, error) {
	return EthAccount{
		Address: hex.EncodeToString(publicAddress),
		ChainID: chainID, PassPhrase: string(passPhrase)}, nil
}

func (signer EthAccount) GetPublicKey() []byte {
	b, _ := hex.DecodeString(signer.Address)
	return b
}

func (signer EthAccount) Sign(b []byte, ctxVars ...interface{}) (uiptypes.Signature, error) {
	// todo: sign b
	return signaturer.FromRaw(b, uiptypes.SignatureTypeUnderlyingType(signaturetype.Secp256k1)), nil
}

func (vc *VesClient) sendAck(acc *uiprpc_base.Account, sessionID []byte, host string, signature []byte) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		vc.logger.Error("did not connect", "error", err)
		return err
	}
	defer conn.Close()
	c := uiprpc.NewVESClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.SessionAckForInit(
		ctx,
		&uiprpc.SessionAckForInitRequest{
			SessionId: sessionID,
			User:      acc,
			UserSignature: &uiprpc_base.Signature{
				SignatureType: 123456,
				Content:       signature,
			},
		})
	if err != nil {
		vc.logger.Error("could not send ack", "error", err)
		return err
	}
	vc.logger.Info("Session ack", "ok", r.GetOk(), "session id", sessionID)
	return nil
}

func (vc *VesClient) informAttestation(grpcHost string, sendingAtte *wsrpc.AttestationReceiveRequest) {
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		vc.logger.Error("VesClient.informAttestation.grpc.Dial.Failed\n",
			"tid", sendingAtte.GetAtte().Tid, "aid", sendingAtte.GetAtte().Aid, "error", err)
		return
	}
	defer conn.Close()

	c := uiprpc.NewVESClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.InformAttestation(
		ctx,
		&uiprpc.AttestationReceiveRequest{
			SessionId: sendingAtte.SessionId,
			Atte:      sendingAtte.Atte,
		},
	)
	if err != nil {
		vc.logger.Error("VesClient.informAttestation.grpc.Send.Failed\n",
			"tid", sendingAtte.GetAtte().Tid, "aid", sendingAtte.GetAtte().Aid, "error", err)
		return
	}

	if !r.GetOk() {
		vc.logger.Error("VesClient.informAttestation.grpc.Result.Failed\n",
			"tid", sendingAtte.GetAtte().Tid, "aid", sendingAtte.GetAtte().Aid)
	}
	return
}
