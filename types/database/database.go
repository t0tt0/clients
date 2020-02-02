package vesdb

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/types"
)

type Database struct {
	sindb          types.Index
	muldb          types.MultiIndex
	sesdb          types.SessionBase
	kvdb           types.SessionKVBase
	storageHandler types.StorageHandler
	dns            types.ChainDNS
}

func (db *Database) InsertAccount(userName string, acc uiptypes.Account) error {
	panic("implement me")
}

func (db *Database) FindUser(userName string) (*model.User, error) {
	panic("implement me")
}

func (db *Database) FindAccounts(userName string, chainID uint64) ([]uiptypes.Account, error) {
	panic("implement me")
}

func (db *Database) HasAccount(userName string, acc uiptypes.Account) (has bool, err error) {
	panic("implement me")
}

func (db *Database) InvertFind(uiptypes.Account) (*model.User, error) {
	panic("implement me")
}

func (db *Database) GetTransactionProof(chainID uiptypes.ChainID, blockID uiptypes.BlockID, color []byte) (uiptypes.MerkleProof, error) {
	panic("implement me")
}

func (db *Database) GetStorageAt(chainID uiptypes.ChainID, typeID uiptypes.TypeID, contractAddress uiptypes.ContractAddress, pos []byte, description []byte) (uiptypes.Variable, error) {
	return db.storageHandler.GetStorageAt(db.sindb, chainID, typeID, contractAddress, pos, description)
}

func (db *Database) SetStorageOf(chainID uiptypes.ChainID, typeID uiptypes.TypeID, contractAddress uiptypes.ContractAddress, pos []byte, description []byte, variable uiptypes.Variable) error {
	return db.storageHandler.SetStorageOf(db.sindb, chainID, typeID, contractAddress, pos, description, variable)
}

func (db *Database) SetChainDNS(dns types.ChainDNS) bool {
	db.dns = dns
	return true
}

func (db *Database) SetSessionKVBase(logicDB types.SessionKVBase) bool {
	db.kvdb = logicDB
	return true
}

func (db *Database) SetIndex(phyDB types.Index) bool {
	db.sindb = phyDB
	return true
}

func (db *Database) SetMultiIndex(phyDB types.MultiIndex) bool {
	db.muldb = phyDB
	return true
}

func (db *Database) SetSessionBase(logicDB types.SessionBase) bool {
	db.sesdb = logicDB
	return true
}


func (db *Database) SetStorageHandler(logicDB types.StorageHandler) bool {
	db.storageHandler = logicDB
	return true
}

func (db *Database) InsertSessionInfo(session types.Session) error {
	return db.sesdb.InsertSessionInfo(db.muldb, db.sindb, session)
}

func (db *Database) FindSessionInfo(isc_address []byte) (types.Session, error) {
	return db.sesdb.FindSessionInfo(db.muldb, db.sindb, isc_address)
}

func (db *Database) UpdateSessionInfo(session types.Session) error {
	return db.sesdb.UpdateSessionInfo(db.muldb, db.sindb, session)
}

func (db *Database) DeleteSessionInfo(isc_address []byte) error {
	return db.sesdb.DeleteSessionInfo(db.muldb, db.sindb, isc_address)
}

func (db *Database) FindTransaction(isc_address []byte, transaction_id uint64, getter func([]byte) error) error {
	return db.sesdb.FindTransaction(db.sindb, isc_address, transaction_id, getter)
}

func (db *Database) ActivateSession(isc_address []byte) {
	db.sesdb.ActivateSession(isc_address)
}

func (db *Database) InactivateSession(isc_address []byte) {
	db.sesdb.InactivateSession(isc_address)
}

func (db *Database) GetChainInfo(chainId uiptypes.ChainID) (types.ChainInfo, error) {
	return db.dns.GetChainInfo(db.sindb, chainId)
}

func (db *Database) SetKV(isc_address, k, v []byte) error {
	return db.kvdb.SetKV(db.sindb, isc_address, k, v)
}

func (db *Database) GetKV(isc_address, k []byte) ([]byte, error) {
	return db.kvdb.GetKV(db.sindb, isc_address, k)
}

func (db *Database) GetGetter(isc_address []byte) types.KVGetter {
	return db.kvdb.GetGetter(db.sindb, isc_address)
}

func (db *Database) GetSetter(isc_address []byte) types.KVSetter {
	return db.kvdb.GetSetter(db.sindb, isc_address)
}
