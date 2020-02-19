package control

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
)

type SessionFSetI interface {
	AckForInit(*model.Session, uiptypes.Account, uiptypes.Signature) error
	FindTransaction([]uint8, int64) ([]uint8, error)
	GetAccounts(*model.Session) ([]uiptypes.Account, error)
	GetAckCount(*model.Session) (int64, error)
	GetTransactingTransaction(*model.Session) ([]uint8, error)
	InitSessionInfo([]uint8, []*uiptypes.TransactionIntent, []*model.SessionAccount) (*model.Session, error)
	NotifyAttestation(*model.Session, types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) error
	ProcessAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error)
	SyncFromISC() error
}
