package control

import (
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type InitializerI interface {
	InitContent(*opintent.RawIntent, []uint8) error
	InitContents([][]uint8) (*opintent.RawIntents, error)
	InitDependencies([][]uint8) (*opintent.RawDependenciesInfo, error)
	Parse(uip.OpIntents) (opintent.TxIntents, error)
	ParseDependencies(opintent.RawDependenciesI, map[string]int) (*opintent.DependenciesInfo, error)
	ParseIntent(opintent.RawIntentI) ([]uip.TxIntentI, error)
	ParseIntents(opintent.RawIntentsI) (opintent.TxIntentsImpl, error)
	TopologicalSort(opintent.ArrayI, []opintent.Dependency) error
}


type BlockChainInterfaceI interface {
	CheckAddress([]uint8) error
	Deserialize([]uint8) (uip.RawTransaction, error)
	GetStorageAt(uint64, uip.TypeID, []uint8, []uint8, []uint8) (uip.Variable, error)
	GetTransactionProof(uint64, []uint8, []uint8) (uip.MerkleProof, error)
	MustWithSigner() bool
	RouteRaw(uint64, uip.RawTransaction) ([]uint8, error)
	RouteWithSigner(uip.Signer) (uip.Router, error)
	Translate(*uip.TransactionIntent, uip.Storage) (uip.RawTransaction, error)
	WaitForTransact(uint64, []uint8, ...interface{}) ([]uint8, []uint8, error)
	ParseTransactionIntent(intent uip.TxIntentI) (uip.TxIntentI, error)
}


