package control

import (
	"encoding/json"

	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/Myriad-Dreamin/go-ves/types"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/db-layer"
	"github.com/jinzhu/gorm"
)

type SessionDBI interface {
	Filter(*gorm_crud_dao.Filter) ([]dblayer.Session, error)
	FilterI(interface{}) (interface{}, error)
	ID(uint) (*dblayer.Session, error)
	ID_(*gorm.DB, uint) (*dblayer.Session, error)
	QueryChain() *dblayer.SessionQuery
	QueryGUID(string) (*dblayer.Session, error)
}

type SessionAccountDBI interface {
	Filter(*gorm_crud_dao.Filter) ([]dblayer.SessionAccount, error)
	FilterI(interface{}) (interface{}, error)
	GetAcknowledged(string) (int64, error)
	GetTotal(string) (int64, error)
	ID(string) ([]dblayer.SessionAccount, error)
	ID_(*gorm.DB, string) (*dblayer.SessionAccount, error)
	QueryChain() *dblayer.SessionAccountQuery
}

type SessionFSetI interface {
	AckForInit(*dblayer.Session, uiptypes.Account, uiptypes.Signature) error
	FindTransaction([]uint8, int64) ([]uint8, error)
	GetAccounts(*dblayer.Session) ([]uiptypes.Account, error)
	GetAckCount(*dblayer.Session) (int64, error)
	GetTransactingTransaction(*dblayer.Session) ([]uint8, error)
	InitSessionInfo([]uint8, []*uiptypes.TransactionIntent, []*dblayer.SessionAccount) (*dblayer.Session, error)
	NotifyAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error)
	ProcessAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error)
	SyncFromISC() error
}

type OpIntentInitializerI interface {
	InitContractInvocationOpIntent(string, json.RawMessage) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	InitOpIntent(uiptypes.OpIntents) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	InitPaymentOpIntent(string, json.RawMessage) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	TopologicalSort([][]*uiptypes.TransactionIntent, []opintent.Dependency) error
}
