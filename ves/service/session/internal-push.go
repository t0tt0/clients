package sessionservice

import (
	"context"
	"encoding/hex"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"time"
)

func (svc *Service) pushTransaction(
	ctx context.Context, ses *model.Session, transactionID int64) (err error) {
	ti, err := svc.getTransactionIntent(ses.GetGUID(), transactionID)
	if err != nil {
		return err
	}
	var accounts []*uiprpc_base.Account
	accounts = append(accounts, &uiprpc_base.Account{
		Address: ti.Src,
		ChainId: ti.ChainID,
	})
	svc.logger.Info("sending attestation request", "chain id", ti.ChainID, "address", hex.EncodeToString(ti.Src))

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	reply, err := svc.cVes.InternalAttestationSending(ctx, &uiprpc.InternalRequestComingRequest{
		SessionId: ses.GetGUID(),
		Host:      svc.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  accounts,
	})
	if err != nil {
		return wrapper.Wrap(types.CodeAttestationSendError, err)
	}
	if reply.GetOk() != true {
		return wrapper.WrapCode(types.CodeAttestationSendError)
	}

	return nil
}

func (svc *Service) pushInternalInitRequestBySessionAccount(ctx context.Context, iscAddress []byte, accounts []*model.SessionAccount) error {
	return svc.pushInternalInitRequest(ctx, iscAddress, func() (uaccs []*uiprpc_base.Account) {
		for _, acc := range accounts {
			uaccs = append(uaccs, &uiprpc_base.Account{
				Address: acc.GetAddress(),
				ChainId: acc.GetChainId(),
			})
		}
		return
	}())
}

func (svc *Service) pushInternalInitRequest(ctx context.Context, iscAddress []byte, accounts []*uiprpc_base.Account) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	r, err := svc.cVes.InternalRequestComing(ctx, &uiprpc.InternalRequestComingRequest{
		SessionId: iscAddress,
		Host:      svc.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  accounts,
	})

	if err != nil {
		return wrapper.Wrap(types.CodeSessionInitInternalRequestError, err)
	}
	if !r.GetOk() {
		return wrapper.WrapCode(types.CodeSessionInitInternalRequestError)
	}
	return nil
}

func (svc *Service) pushInternalCloseRequestBySessionAccount(ctx context.Context, iscAddress []byte, accounts []model.SessionAccount) error {
	return svc.pushInternalCloseRequest(ctx, iscAddress, func() (uaccs []*uiprpc_base.Account) {
		for _, acc := range accounts {
			uaccs = append(uaccs, &uiprpc_base.Account{
				Address: acc.GetAddress(),
				ChainId: acc.GetChainId(),
			})
		}
		return
	}())
}

func (svc *Service) pushInternalCloseRequest(ctx context.Context, iscAddress []byte, accounts []*uiprpc_base.Account) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := svc.cVes.InternalCloseSession(ctx, &uiprpc.InternalCloseSessionRequest{
		SessionId: iscAddress,
		NsbHost:   svc.cfg.BaseParametersConfig.NSBHost,
		GrpcHost:  svc.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  accounts,
	})
	if err != nil {
		return wrapper.Wrap(types.CodeSessionCloseInternalRequestError, err)
	}
	if r.GetOk() != true {
		return wrapper.WrapCode(types.CodeAttestationSendError)
	}

	return nil
}
