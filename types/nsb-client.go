package types

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type NSBClient interface {
	FreezeInfo(signer uiptypes.Signer, guid []byte, u uint64) ([]byte, error)
	AddMerkleProof(signer uiptypes.Signer, nil2 interface{}, underlyingType uiptypes.MerkleProofTypeUnderlyingType, hash uiptypes.RootHash, proof uiptypes.Proof, key uiptypes.MerkleProofKey, value uiptypes.MerkleProofValue) ([]byte, error)

}


