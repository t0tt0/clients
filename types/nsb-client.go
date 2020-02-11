package types

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	nsb_message "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client/nsb-message"
)

type NSBClient interface {
	FreezeInfo(signer uiptypes.Signer, guid []byte, u uint64) ([]byte, error)
	AddMerkleProof(user uiptypes.Signer, toAddress []byte,
		merkleType uint16, rootHash, proof, key, value []byte, ) (*nsb_message.ResultInfo, error)
	AddBlockCheck(
		user uiptypes.Signer, toAddress []byte,
		chainID uint64, blockID, rootHash []byte, rcType uint8,
	) (*nsb_message.ResultInfo, error)
	InsuranceClaim(
		user uiptypes.Signer, contractAddress []byte,
		tid, aid uint64,
	) (*nsb_message.DeliverTx, error)
}
