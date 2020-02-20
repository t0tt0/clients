package control

import "github.com/HyperService-Consortium/go-uip/uip"

type SessionKV interface {
	SetKV(iscAddress []byte, provedKey []byte, provedValue []byte) error
	GetKV(iscAddress []byte, provedKey []byte) (provedValue []byte, err error)
}

type StorageHandler interface {
	GetTransactionProof(chainID uip.ChainID, blockID uip.BlockID, color []byte) (uip.MerkleProof, error)
	GetStorageAt(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error)
	SetStorageOf(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte, variable uip.Variable) error
}
