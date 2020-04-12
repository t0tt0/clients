package dep_uip

import (
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type TransactionIntent struct{ *opintent.TransactionIntent }
type ContractInvokeMeta = lexer.ContractInvokeMeta

func (t TransactionIntent) GetTxType() trans_type.Type {
	return t.TransType
}

func (t TransactionIntent) GetChainID() uip.ChainIDUnderlyingType {
	return t.ChainID
}

func DecTransactionIntent(intent *opintent.TransactionIntent) *TransactionIntent {
	return &TransactionIntent{TransactionIntent: intent}
}
