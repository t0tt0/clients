package control

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type SessionKV interface {
	SetKV(iscAddress []byte, provedKey []byte, provedValue []byte) error
	GetKV(iscAddress []byte, provedKey []byte) (provedValue []byte, err error)
}

type StorageHandler interface {
	GetTransactionProof(chainID uiptypes.ChainID, blockID uiptypes.BlockID, color []byte) (uiptypes.MerkleProof, error)
	GetStorageAt(chainID uiptypes.ChainID, typeID uiptypes.TypeID, contractAddress uiptypes.ContractAddress, pos []byte, description []byte) (uiptypes.Variable, error)
	SetStorageOf(chainID uiptypes.ChainID, typeID uiptypes.TypeID, contractAddress uiptypes.ContractAddress, pos []byte, description []byte, variable uiptypes.Variable) error
}
