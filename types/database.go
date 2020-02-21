package types

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

type KVPair struct {
	Key   string
	Value interface{}
}

type KVMap = map[string]interface{}

type KVObject interface {
	GetObjectPtr() interface{}
	GetSlicePtr() interface{}
	GetID() int64
	ToKVMap() KVMap
}

type Index interface {
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	Delete([]byte) error
	Batch([][]byte, [][]byte) error
}

type MultiIndex interface {
	// RegisterObject(KVObject) error

	Insert(KVObject) error

	Get(KVObject) (bool, error)

	Select(KVObject) (interface{}, error)

	SelectAll(KVObject) (interface{}, error)

	// 要求只Delete到一个
	Delete(KVObject) error

	// 可以Delete多个
	MultiDelete(KVObject) error

	Modify(KVObject, KVMap) error

	MultiModify(KVObject, KVMap) error
}

type KVPMultiIndex interface {
	Insert(...KVPair) error

	Select([]interface{}, ...KVPair) error

	// 要求只Delete到一个
	Delete(...KVPair) error
	// 可以Delete多个
	MultiDelete(...KVPair) error

	// 要求只Update到一个
	Modify([]KVPair, ...KVPair) error
	// 可以Update到多个
	MultiModify([]KVPair, ...KVPair) error
}

type ORMMultiIndex interface {
	MultiIndex
	// 要求只Update到一个
	// Modify(ORMObject, ORMObject) error
	// 可以Update到多个
	// MultiModify(ORMObject, ORMObject) error
}

type chain_id = uint64

type ChainInfo interface {
	GetChainType() uip.ChainType
	GetChainHost() string
}

type StorageHandlerInterface interface {
	GetTransactionProof(chainID uip.ChainID, blockID uip.BlockID, color []byte) (uip.MerkleProof, error)
	GetStorageAt(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error)
	SetStorageOf(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte, variable uip.Variable) error
}
