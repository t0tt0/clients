package sessionservice

import (
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"golang.org/x/net/context"

	transtype "github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
)

func (svc *Service) SessionRequireRawTransact(
	ctx context.Context, in *uiprpc.SessionRequireRawTransactRequest) (
	*uiprpc.SessionRequireRawTransactReply, error) {

	ses, err := svc.getSession(in.GetSessionId())
	if err != nil {
		return nil, err
	}

	ti, bn, err := svc.getCurrentTxIntentWithExecutor(ses)
	if err != nil {
		return nil, err
	}

	if err = newPrepareTranslateEnvironment(svc, ses, ti, bn).do(); err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionPrepareTranslateError, err)
	}

	var b uip.RawTransaction
	b, err = bn.Translate(ti, svc.storageHandler)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionTranslateError, err)
	}

	x, err := b.Serialize()
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionRawSerializeError, err)
	}

	var dest *uiprpc_base.Account
	switch ti.TransType {
	case transtype.Payment:
		dest = &uiprpc_base.Account{
			Address: ti.Dst,
			ChainId: ti.ChainID,
		}
	case transtype.ContractInvoke:
		dest = &uiprpc_base.Account{
			Address: svc.respAccount.GetAddress(),
			ChainId: svc.respAccount.GetChainId(),
		}
	default:
		return nil, wrapper.WrapCode(types.CodeDestinationRespUnknown)
	}

	return &uiprpc.SessionRequireRawTransactReply{
		RawTransaction: x,
		Tid:            uint64(ses.UnderTransacting),
		Src: &uiprpc_base.Account{
			Address: ti.Src,
			ChainId: ti.ChainID,
		},
		Dst: dest,
	}, nil

}
