package sessionservice

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model"
)

func (svc *Service) getCurrentTxIntentWithExecutor(ses *model.Session) (*opintent.TransactionIntent, uip.BlockChainInterface, error) {
	ti, err := svc.getTransactionIntent(ses.GetGUID(), ses.UnderTransacting)
	if err != nil {
		return nil, nil, wrapper.Wrap(types.CodeGetTransactionIntentError, err)
	}
	if ti.GetType() == trans_type.Payment || ti.GetType() == trans_type.ContractInvoke {
		ti := ti.(*opintent.TransactionIntent)
		bn, err := svc.getBlockChainInterface(ti.ChainID)
		if err != nil {
			return nil, nil, wrapper.Wrap(types.CodeGetBlockChainInterfaceError, err)
		}
		return ti, bn, nil
	} else {
		return nil, nil, fmt.Errorf("gotten type = %v, which is not a valid transaction", ti.GetType())
	}
}
