package fset

import (
	"fmt"
	"github.com/HyperService-Consortium/NSB/contract/isc/TxState"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/lib/basic/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/upstream"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	database2 "github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
)

// SessionFSet is the collection of functions related to model.session
type SessionFSet struct {
	AccountDB     abstraction.SessionAccountDB
	Index         types.Index
	TransactionDB abstraction.TransactionDB
	SessionDB     abstraction.SessionDB
}

func NewSessionFSet(provider abstraction.Provider, index types.Index) *SessionFSet {
	return &SessionFSet{
		TransactionDB: provider.TransactionDB(),
		AccountDB:     provider.SessionAccountDB(),
		SessionDB:     provider.SessionDB(),
		Index:         index,
	}
}

func (s SessionFSet) GetAccounts(ses *database2.Session) ([]uiptypes.Account, error) {
	accounts, err := s.AccountDB.ID(ses.ISCAddress)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAccountFindError, err)
	}
	return database2.SessionAccountsToUIPAccounts(accounts), nil
}

func (s SessionFSet) GetAckCount(ses *database2.Session) (int64, error) {
	return s.AccountDB.GetAcknowledged(ses.ISCAddress)
}

func (s SessionFSet) FindTransaction(
	iscAddress []byte,
	transactionID int64) (b []byte, err error) {
	tx := new(database2.Transaction)
	has, err := s.TransactionDB.FindSessionIndex(tx, database2.EncodeAddress(iscAddress), transactionID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSelectError, err)
	} else if !has {
		return nil, wrapper.WrapCode(types.CodeNotFound)
	}

	return database2.DecodeContent(tx.Content)
}

//?
func (s SessionFSet) GetTransactingTransaction(ses *database2.Session) ([]byte, error) {
	//return ses.UnderTransacting
	return s.FindTransaction(ses.GetGUID(), ses.UnderTransacting)
}

func (s SessionFSet) InitSessionInfo(
	iscAddress []byte, intents []*opintent.TransactionIntent, accounts []*database2.SessionAccount) (
	ses *database2.Session, err error) {
	ses = database2.NewSession(iscAddress)
	for i := range accounts {
		accounts[i].SessionID = ses.ISCAddress
		if _, err = s.AccountDB.Create(accounts[i]); err != nil {
			err = wrapper.Wrap(types.CodeSessionInsertAccountError, err)
			return
		}
	}

	ses.AccountsCount, err = s.AccountDB.GetTotal(ses.ISCAddress)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAccountGetTotalError, err)
	}

	ses.TransactionCount = int64(len(intents))
	for i := range intents {
		b, err := upstream.Serializer.TransactionIntent.Marshal(intents[i])
		if err != nil {
			return nil, wrapper.Wrap(types.CodeTransactionIntentSerializeError, err)
		}

		if _, err = s.TransactionDB.Create(&database2.Transaction{
			SessionID: ses.ISCAddress,
			Index:     int64(i),
			Content:   database2.EncodeContent(b),
		}); err != nil {
			return nil, wrapper.Wrap(types.CodeSessionInsertTransactionError, err)
		}
	}

	_, err = s.SessionDB.Create(ses)
	return
}

func (s SessionFSet) AckForInit(ses *database2.Session, acc uiptypes.Account, signature uiptypes.Signature) error {
	var (
		sac = &database2.SessionAccount{
			SessionID: ses.ISCAddress,
			ChainID:   acc.GetChainId(),
			Address:   encoding.EncodeBase64(acc.GetAddress()),
		}
		has bool
		err error
	)
	//verifySignature

	if has, err = s.AccountDB.Find(sac); err != nil {
		return wrapper.Wrap(types.CodeSessionAccountFindError, err)
	} else if !has {
		return wrapper.Wrap(types.CodeSessionAccountNotFound, err)
	}

	sac.Acknowledged = true
	if aff, err := s.AccountDB.UpdateAcknowledged(sac); err != nil {
		return wrapper.Wrap(types.CodeUpdateError, err)
	} else if aff == 0 {
		return wrapper.WrapCode(types.CodeUpdateNoEffect)
	}
	return nil
}

type NotCurrentTransactionError struct {
	Requiring int64 `json:"requiring"`
	Current   int64 `json:"current"`
}

func (n NotCurrentTransactionError) Error() string {
	return fmt.Sprintf("requiring:%v,current:%v", n.Requiring, n.Current)
}

func (s SessionFSet) NotifyAttestation(
	ses *database2.Session, nsb types.NSBInterface,
	bn uiptypes.BlockChainInterface, attestation uiptypes.Attestation) (err error) {
	if attestation.GetTid() != uint64(ses.UnderTransacting) {
		return wrapper.Wrap(types.CodeSessionNotCurrentTransaction, NotCurrentTransactionError{
			Requiring: int64(attestation.GetTid()),
			Current:   ses.UnderTransacting,
		})
	}
	switch attestation.GetAid() {
	case TxState.Unknown, TxState.Initing, TxState.Inited:
		return nil
	case TxState.Instantiating, TxState.Instantiated, TxState.Open, TxState.Opened:
		return nil
	case TxState.Closed:
		ses.UnderTransacting++
		if ses.UnderTransacting == ses.TransactionCount {
			err = nsb.SettleContract(ses.GetGUID())
			if err != nil {
				return wrapper.Wrap(types.CodeSettleContractError, err)
			}
		}
		if _, err = s.SessionDB.Update(ses); err != nil {
			return wrapper.Wrap(types.CodeUpdateError, err)
		}

		return nil
	default:
		return wrapper.WrapCode(types.CodeTransactionStateNotFound)
	}
}

func (s SessionFSet) ProcessAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error) {
	panic("implement me")
}

func (s SessionFSet) SyncFromISC() error {
	panic("implement me")
}
