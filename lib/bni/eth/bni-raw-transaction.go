package bni

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/go-uip/uip"
	ethclient "github.com/HyperService-Consortium/go-ves/lib/net/eth-client"
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
	uip.Signer
	GetEthPassword() string
}

type passwordSigner struct {
	pb []byte
	ps string
}

func (p passwordSigner) GetPublicKey() uip.PublicKey {
	return p.pb
}

var ErrNotUnlock = errors.New("error not unlock")

func (p passwordSigner) Sign(content uip.SignatureContent, ctxVars ...interface{}) (uip.Signature, error) {

	var (
		duration  int
		chainInfo uip.ChainInfo
	)

	for _, rawV := range ctxVars {
		switch v := rawV.(type) {
		case uip.SignerOptionDuration:
			duration = int(v) / 1000
		case uip.SignerOptionChainInfo:
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
func (t RawTransaction) Sign(signer uip.Signer, ctxVars ...interface{}) (uip.RawTransaction, error) {
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
