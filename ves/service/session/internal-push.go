package objectservice

import (
	"context"
	"encoding/hex"
	"encoding/json"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"time"
)

func (svc *Service) pushTransaction(
	ctx context.Context, ses *model.Session, transactionID int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	txb, err := svc.sesFSet.FindTransaction(ses.GetGUID(), transactionID)
	if err != nil {
		return wrapper.Wrap(types.CodeTransactionFindError, err)
	}
	var kvs opintent.TransactionIntent
	err = json.Unmarshal(txb, &kvs)
	if err != nil {
		return wrapper.Wrap(types.CodeDeserializeTransactionError, err)
	}
	var accounts []*uiprpc_base.Account
	accounts = append(accounts, &uiprpc_base.Account{
		Address: kvs.Src,
		ChainId: kvs.ChainID,
	})
	svc.logger.Info("sending attestation request", "chain id", kvs.ChainID, "address", hex.EncodeToString(kvs.Src))

	_, err = svc.cVes.InternalAttestationSending(ctx, &uiprpc.InternalRequestComingRequest{
		SessionId: ses.GetGUID(),
		Host:      svc.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  accounts,
	})
	if err != nil {
		return wrapper.Wrap(types.CodeAttestationSendError, err)
	}
	return nil
}
func (svc *Service) pushInternalInitRequest(ctx context.Context, iscAddress []byte, accounts []*model.SessionAccount) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	r, err := svc.cVes.InternalRequestComing(ctx, &uiprpc.InternalRequestComingRequest{
		SessionId: iscAddress,
		Host:      svc.cfg.BaseParametersConfig.ExposeHost,
		Accounts: func() (uaccs []*uiprpc_base.Account) {
			for _, acc := range accounts {
				uaccs = append(uaccs, &uiprpc_base.Account{
					Address: acc.GetAddress(),
					ChainId: acc.GetChainId(),
				})
			}
			return
		}(),
	})

	if err != nil {
		return false, wrapper.Wrap(types.CodeSessionInitInternalRequestError, err)
	}
	return r.Ok, nil
}
