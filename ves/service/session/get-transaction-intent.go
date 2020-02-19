package sessionservice

import (
	"encoding/json"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
)

func (svc *Service) getTransactionIntent(sessionID []byte, transactionID int64) (*opintent.TransactionIntent, error) {
	txb, err := svc.sesFSet.FindTransaction(sessionID, transactionID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionFindError, err)
	}
	var ti opintent.TransactionIntent
	err = json.Unmarshal(txb, &ti)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeDeserializeTransactionError, err)
	}
	return &ti, nil
}
