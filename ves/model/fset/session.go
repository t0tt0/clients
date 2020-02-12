package fset

import (
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/config"
	bitmap "github.com/Myriad-Dreamin/go-ves/lib/bitmapping/redis-bitmap"
	mredis "github.com/Myriad-Dreamin/go-ves/lib/database/redis"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/serial_helper"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
)

// SessionFSet is the collection of functions related to model.session
type SessionFSet struct {
	AccountDB *model.SessionAccountDB
	Index     types.Index
}

func NewSessionFSet(provider *model.Provider, index types.Index) *SessionFSet {
	return &SessionFSet{AccountDB: provider.SessionAccountDB(), Index: index}
}

func (s SessionFSet) GetAccounts(ses *model.Session) ([]uiptypes.Account, error) {
	accounts, err := s.AccountDB.ID(ses.ISCAddress)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAccountFindError, err)
	}
	return model.SessionAccountsToUIPAccounts(accounts), nil
}

//func (s SessionFSet) GetTransaction(interface{}) interface{} {
//	panic("implement me")
//}
//
//func (s SessionFSet) GetTransactions() []interface{} {
//	panic("implement me")
//}

func (s SessionFSet) GetAckCount(ses *model.Session) (int64, error) {
	count, err := bitmap.GetBitMap(ses.GetGUID(), mredis.RedisCacheClient.Pool.Get()).Count()
	if err != nil {
		return 0, wrapper.Wrap(types.CodeSessionRedisGetAckCountError, err)
	}
	return count, nil
}

func (s SessionFSet) FindTransaction(
	iscAddress []byte,
	transactionID int64) (b []byte, err error) {
	var k []byte
	k, err = serial_helper.DecoratePrefix([]byte{
		uint8((transactionID >> 56) & 0xff), uint8((transactionID >> 48) & 0xff),
		uint8((transactionID >> 40) & 0xff), uint8((transactionID >> 32) & 0xff),
		uint8((transactionID >> 24) & 0xff), uint8((transactionID >> 16) & 0xff),
		uint8((transactionID >> 8) & 0xff), uint8((transactionID >> 0) & 0xff),
	}, iscAddress)
	if err != nil {
		return
	}
	k, err = serial_helper.DecoratePrefix(config.TransactionPrefix, k)
	if err != nil {
		return
	}
	b, err = s.Index.Get(k)
	if err != nil {
		return
	}
	return
}

//?
func (s SessionFSet) GetTransactingTransaction(ses *model.Session) ([]byte, error) {
	//return ses.UnderTransacting
	return s.FindTransaction(ses.GetGUID(), ses.UnderTransacting)
}

func (s SessionFSet) InitSessionInfo(
	iscAddress []byte, intents []*opintent.TransactionIntent, accounts []*model.SessionAccount) (
	ses *model.Session, err error) {
	ses = model.NewSession(iscAddress)
	ses.AccountsCount, err = s.AccountDB.GetTotal(ses.ISCAddress)
	if err != nil {
		err = wrapper.Wrap(types.CodeSessionAccountGetTotolError, err)
		return
	}
	for i := range accounts {
		accounts[i].SessionID = ses.ISCAddress
		if _, err = accounts[i].Create(); err != nil {
			err = wrapper.Wrap(types.CodeSessionInsertAccountError, err)
			return
		}
	}

	//todo add transaction
	//_ = transactions

	_, err = ses.Create()
	return
}

func (s SessionFSet) AckForInit(ses *model.Session, acc uiptypes.Account, signature uiptypes.Signature) error {
	var (
		//addr         = acc.GetAddress()
		//acknowledges = bitmap.GetBitMap(
		//	ses.GetGUID(), mredis.RedisCacheClient.Pool.Get())
		sac = &model.SessionAccount{
			SessionID: ses.ISCAddress,
			ChainID:   acc.GetChainId(),
			Address:   encoding.EncodeBase64(acc.GetAddress()),
		}
		has bool
		err error
	)
	//todo: database transaction
	//verifySignature

	if has, err = sac.Find(); err != nil {
		return wrapper.Wrap(types.CodeSessionAccountFindError, err)
	} else if !has {
		return wrapper.Wrap(types.CodeSessionAccountNotFound, err)
	}

	sac.Acknowledged = true
	if aff, err := sac.Update(); err != nil {
		return wrapper.Wrap(types.CodeUpdateError, err)
	} else if aff == 0 {
		return wrapper.WrapCode(types.CodeUpdateNoEffect)
	}
	// todo: NSB
	return nil
}

func (s SessionFSet) NotifyAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error) {
	panic("implement me")
}

func (s SessionFSet) ProcessAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error) {
	panic("implement me")
}

func (s SessionFSet) SyncFromISC() error {
	panic("implement me")
}
