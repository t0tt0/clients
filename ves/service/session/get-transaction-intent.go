package sessionservice

import (
	"bytes"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
)

func (svc *Service) getTransactionIntent(sessionID []byte, transactionID int64) (uip.Instruction, error) {
	txb, err := svc.sesFSet.FindTransaction(sessionID, transactionID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionFindError, err)
	}
	var ti uip.Instruction
	ti, err = opintent.DecodeInstruction(bytes.NewReader(txb))
	if err != nil {
		return nil, wrapper.Wrap(types.CodeDeserializeTransactionError, err)
	}
	return ti, nil
}
