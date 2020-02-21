package types

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	nsb_message "github.com/HyperService-Consortium/go-ves/lib/net/nsb-client/nsb-message"
)

type NSBClient interface {
	FreezeInfo(signer uip.Signer, guid []byte, u uint64) ([]byte, error)
	AddMerkleProof(user uip.Signer, toAddress []byte,
		merkleType uint16, rootHash, proof, key, value []byte) (*nsb_message.ResultInfo, error)
	AddBlockCheck(
		user uip.Signer, toAddress []byte,
		chainID uint64, blockID, rootHash []byte, rcType uint8,
	) (*nsb_message.ResultInfo, error)
	InsuranceClaim(
		user uip.Signer, contractAddress []byte,
		tid, aid uint64,
	) (*nsb_message.DeliverTx, error)
	CreateISC(signer uip.Signer, uint32s []uint32, bytes [][]byte, bytes2 [][]byte, bytes3 []byte) ([]byte, error)
	SettleContract(signer uip.Signer, bytes []byte) (*nsb_message.DeliverTx, error)
}

