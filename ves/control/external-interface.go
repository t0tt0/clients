package control

import (
	"encoding/json"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	nsb_message "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client/nsb-message"
	"github.com/Myriad-Dreamin/go-ves/types"
)

type Signer = uiptypes.Signer
type Logger = types.Logger
type CenteredVESClient = uiprpc.CenteredVESClient
type Account = uiptypes.Account
type SessionKV = types.SessionKV
type StorageHandler = types.StorageHandler
type ChainDNSInterface = types.ChainDNSInterface

type NSBClient interface {
	FreezeInfo(signer uiptypes.Signer, guid []byte, u uint64) ([]byte, error)
	AddMerkleProof(user uiptypes.Signer, toAddress []byte,
		merkleType uint16, rootHash, proof, key, value []byte) (*nsb_message.ResultInfo, error)
	AddBlockCheck(
		user uiptypes.Signer, toAddress []byte,
		chainID uint64, blockID, rootHash []byte, rcType uint8,
	) (*nsb_message.ResultInfo, error)
	InsuranceClaim(
		user uiptypes.Signer, contractAddress []byte,
		tid, aid uint64,
	) (*nsb_message.DeliverTx, error)
	CreateISC(signer uiptypes.Signer, uint32s []uint32, bytes [][]byte, bytes2 [][]byte, bytes3 []byte) ([]byte, error)
}

type OpIntentInitializerI interface {
	InitContractInvocationOpIntent(string, json.RawMessage) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	InitOpIntent(uiptypes.OpIntents) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	InitPaymentOpIntent(string, json.RawMessage) ([]*uiptypes.TransactionIntent, []*uiptypes.MerkleProofProposal, error)
	TopologicalSort([][]*uiptypes.TransactionIntent, []opintent.Dependency) error
}