package control

import (
	"encoding/json"

	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
)

type OpIntentInitializerI interface {
	InitContractInvocationOpIntent(string, json.RawMessage) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	InitOpIntent(uiptypes.OpIntents) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	InitPaymentOpIntent(string, json.RawMessage) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	TopologicalSort([][]*uiptypes.TransactionIntent, []opintent.Dependency) error
}

type BlockChainInterfaceI interface {
	CheckAddress([]uint8) error
	Deserialize([]uint8) (uiptypes.RawTransaction, error)
	GetStorageAt(uint64, uiptypes.TypeID, []uint8, []uint8, []uint8) (uiptypes.Variable, error)
	GetTransactionProof(uint64, []uint8, []uint8) (uiptypes.MerkleProof, error)
	MustWithSigner() bool
	RouteRaw(uint64, uiptypes.RawTransaction) ([]uint8, error)
	RouteWithSigner(uiptypes.Signer) (uiptypes.Router, error)
	Translate(*uiptypes.TransactionIntent, uiptypes.Storage) (uiptypes.RawTransaction, error)
	WaitForTransact(uint64, []uint8, ...interface{}) ([]uint8, []uint8, error)
}
