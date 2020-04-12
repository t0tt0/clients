package control

import (
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type InitializerI interface {
	Parse(opintent.OpIntents) (parser.TxIntents, error)
	ParseR(opintent.OpIntentsPacket) (parser.TxIntents, error)
	Parse_(*parser.LexerResult) (parser.TxIntents, error)
	TopologicalSort(opintent.ArrayI, []parser.Dependency) error
}

type BlockChainInterfaceI interface {
	CheckAddress([]uint8) error
	Deserialize([]uint8) (uip.RawTransaction, error)
	GetStorageAt(uint64, uint16, []uint8, []uint8, []uint8) (uip.Variable, error)
	GetTransactionProof(uint64, []uint8, []uint8) (uip.MerkleProof, error)
	MustWithSigner() bool
	ParseTransactionIntent(uip.TxIntentI) (uip.TxIntentI, error)
	RouteRaw(uint64, uip.RawTransaction) ([]uint8, error)
	RouteWithSigner(uip.Signer) (uip.Router, error)
	Translate(uip.TransactionIntent, uip.Storage) (uip.RawTransaction, error)
	WaitForTransact(uint64, []uint8, ...interface{}) ([]uint8, []uint8, error)
}
