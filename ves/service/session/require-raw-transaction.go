package sessionservice

import (
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"golang.org/x/net/context"

	transtype "github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
)

func (svc *Service) RequireRawTransaction(
	ctx context.Context, in *uiprpc.SessionRequireRawTransactRequest) (
	*uiprpc.SessionRequireRawTransactReply, error) {

	ses, err := svc.db.QueryGUIDByBytes(in.GetSessionId())
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionFindError, err)
	} else if ses == nil {
		return nil, wrapper.WrapCode(types.CodeSessionNotFind)
	}

	ti, err := svc.getTransactionIntent(ses.GetGUID(), ses.UnderTransacting)
	if err != nil {
		return nil, err
	}

	bn, err := svc.getBlockChainInterface(ti.ChainID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeGetBlockChainInterfaceError, err)
	}

	if err = newPrepareTranslateEnvironment(svc, ses, ti, bn).do(); err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionPrepareTranslateError, err)
	}

	var b uiptypes.RawTransaction
	b, err = bn.Translate(ti, svc.storageHandler)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionTranslateError, err)
	}

	x, err := b.Serialize()
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionRawSerializeError, err)
	}

	var dest *uiprpc_base.Account
	if ti.TransType == transtype.Payment {
		dest = &uiprpc_base.Account{
			Address: ti.Dst,
			ChainId: ti.ChainID,
		}
	} else {
		dest = &uiprpc_base.Account{
			Address: svc.respAccount.GetAddress(),
			ChainId: svc.respAccount.GetChainId(),
		}
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
