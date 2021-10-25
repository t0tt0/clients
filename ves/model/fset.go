package model

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/fset"
)

type SessionFSet = fset.SessionFSet

func NewSessionFSet(p Provider, index types.Index) *SessionFSet {
	return fset.NewSessionFSet(p, index)
}

type SessionFSetI interface {
	AckForInit(*Session, uip.Account, uip.Signature) error
	FindTransaction([]uint8, int64) ([]uint8, error)
	GetAccounts(*Session) ([]uip.Account, error)
	GetAckCount(*Session) (int64, error)
	GetTransactingTransaction(*Session) ([]uint8, error)
	InitSessionInfo([]uint8, []uip.Instruction, []*SessionAccount) (*Session, error)
	NotifyAttestation(*Session, types.NSBInterface, uip.BlockChainInterface, uip.Attestation) error
	ProcessAttestation(types.NSBInterface, uip.BlockChainInterface, uip.Attestation) (interface{}, interface{}, error)
	SyncFromISC() error
}
