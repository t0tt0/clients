package sessionservice

import (
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
)

func (svc *Service) getCurrentTxIntentWithExecutor(ses *model.Session) (*opintent.TransactionIntent, uiptypes.BlockChainInterface, error) {
	ti, err := svc.getTransactionIntent(ses.GetGUID(), ses.UnderTransacting)
	if err != nil {
		return nil, nil, wrapper.Wrap(types.CodeGetTransactionIntentError, err)
	}

	bn, err := svc.getBlockChainInterface(ti.ChainID)
	if err != nil {
		return nil, nil, wrapper.Wrap(types.CodeGetBlockChainInterfaceError, err)
	}
	return ti, bn, nil
}
