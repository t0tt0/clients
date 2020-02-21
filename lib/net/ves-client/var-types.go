package vesclient

import (
	"context"
	"encoding/hex"
	signaturetype "github.com/HyperService-Consortium/go-uip/const/signature_type"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	uip "github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
	"google.golang.org/grpc"
	"time"
)

// ECCKey is the private key object in memory
type ECCKey struct {
	PrivateKey []byte                         `json:"private_key"`
	ChainID    uip.ChainIDUnderlyingType `json:"chain_id"`
}

// ECCKeyAlias is the private key object in json
type ECCKeyAlias struct {
	PrivateKey string                         `json:"private_key"`
	ChainID    uip.ChainIDUnderlyingType `json:"chain_id"`
	Alias      string                         `json:"alias"`
}

// EthAccount is the account object in memory
type EthAccount struct {
	Address    string                         `json:"address"`
	ChainID    uip.ChainIDUnderlyingType `json:"chain_id"`
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

func NewEthAccount(publicAddress, passPhrase []byte,
	chainID uip.ChainIDUnderlyingType) (EthAccount, error) {
	return EthAccount{
		Address: hex.EncodeToString(publicAddress),
		ChainID: chainID, PassPhrase: string(passPhrase)}, nil
}

func (signer EthAccount) GetPublicKey() []byte {
	b, _ := hex.DecodeString(signer.Address)
	return b
}

func (signer EthAccount) GetEthPassword() string {
	return signer.PassPhrase
}

func (signer EthAccount) Sign(b []byte, ctxVars ...interface{}) (uip.Signature, error) {
	// todo: sign b
	return signaturer.FromRaw(b, uip.SignatureTypeUnderlyingType(signaturetype.Secp256k1)), nil
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
