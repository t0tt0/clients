package bni

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	ethclient "github.com/Myriad-Dreamin/go-ves/lib/net/eth-client"
)

type RawTransaction struct {
	B           []byte `json:"b" form:"b"`
	Responsible []byte `json:"r" form:"r"`
	IsSigned    bool   `json:"i" form:"i"`
}

func NewRawTransaction(b []byte, responsible []byte, isSigned bool) *RawTransaction {
	return &RawTransaction{B: b, Responsible: responsible, IsSigned: isSigned}
}

func (t RawTransaction) Serialize() ([]byte, error) {
	return json.Marshal(&t)
}

func (t RawTransaction) Bytes() ([]byte, error) {
	return t.B, nil
}

func (t RawTransaction) Signed() bool {
	return t.IsSigned
}

type PasswordSigner interface {
	uiptypes.Signer
	GetEthPassword() string
}

type passwordSigner struct {
	pb []byte
	ps string
}

func (p passwordSigner) GetPublicKey() uiptypes.PublicKey {
	return p.pb
}

var ErrNotUnlock = errors.New("error not unlock")

func (p passwordSigner) Sign(content uiptypes.SignatureContent, ctxVars ...interface{}) (uiptypes.Signature, error) {

	var (
		duration  int
		chainInfo uiptypes.ChainInfo
	)

	for _, rawV := range ctxVars {
		switch v := rawV.(type) {
		case uiptypes.SignerOptionDuration:
			duration = int(v) / 1000
		case uiptypes.SignerOptionChainInfo:
			chainInfo = v
		}
	}
	if duration < 10 {
		duration = 10
	}
	if chainInfo == nil {
		return nil, ErrHasNoChainInfo
	}
	unlock, err := ethclient.NewEthClient(chainInfo.GetChainHost()).
		PersonalUnlockAccout(
			hex.EncodeToString(p.pb), p.ps, duration)
	if err != nil {
		return nil, err
	}
	if !unlock {
		return nil, ErrNotUnlock
	}
	return nil, nil
}

func (p passwordSigner) GetEthPassword() string {
	return p.ps
}

// todo change raw transaction signature = sign(signer, context)
func (t RawTransaction) Sign(signer uiptypes.Signer, ctxVars ...interface{}) (uiptypes.RawTransaction, error) {
	address := signer.GetPublicKey()
	if !bytes.Equal(address, t.Responsible) {
		return t, ErrNotMatchAddress
	}

	var err error
	if _, ok := signer.(PasswordSigner); ok {
		_, err = signer.Sign(t.B, ctxVars...)
	} else {
		panic("todo")
	}

	t.IsSigned = true
	return t, err
}
