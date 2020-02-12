package objectservice

import (
	"bytes"
	"context"
	"encoding/hex"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/lib/uniquer"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
)

func (svc *Service) SessionStart(ctx context.Context, in *uiprpc.SessionStartRequest) (*uiprpc.SessionStartReply, error) {
	var (
		ses        *model.Session
		intents    []*opintent.TransactionIntent
		accounts   []*model.SessionAccount
		iscAddress []byte
		ok         bool
		err        error
	)

	if intents, accounts, err = svc.initOpIntents(in.GetOpintents()); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitOpIntentsError, err)
	}
	if iscAddress, err = svc.initISCAddress(intents, accounts); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitGUIDError, err)
	}
	if ses, err = svc.sesFSet.InitSessionInfo(iscAddress, intents, accounts); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitError, err)
	}
	svc.logger.Info("new session requested", "address", hex.EncodeToString(ses.GetGUID()))

	// initializing accounts' bitmap in redis here a long time ago
	//if err = ses.InitAccountRedisMap(); err != nil {
	//	return nil, wrapper.Wrap(types.CodeSessionInitGUIDError, err)
	//}
	//

	for i := range intents {
		//s.Logger.Info()
		if _, err = svc.nsbClient.FreezeInfo(svc.signer, ses.GetGUID(), uint64(i)); err != nil {
			return nil, wrapper.Wrap(types.CodeSessionFreezeInfoError, err)
		}
	}

	if ok, err = svc.pushInternalInitRequest(ctx, iscAddress, accounts); err != nil {
		return nil, err
	}

	return &uiprpc.SessionStartReply{
		Ok:        ok,
		SessionId: iscAddress,
	}, nil
}

func (svc *Service) initOpIntents(opIntents uiptypes.OpIntents) (
	intents []*opintent.TransactionIntent, accounts []*model.SessionAccount, err error) {
	intents, _, err = svc.opInitializer.InitOpIntent(opIntents)
	if err != nil {
		return
	}
	c := uniquer.MakeUniquer()
	if c.Insert(svc.respAccount.GetChainId(), svc.respAccount.GetAddress()) {
		accounts = append(accounts, model.NewSessionAccount(svc.respAccount.GetChainId(), svc.respAccount.GetAddress()))
	}
	for _, intent := range intents {
		//transactions = append(transactions, intent.Bytes())
		if c.Insert(intent.ChainID, intent.Src) {
			accounts = append(accounts, model.NewSessionAccount(intent.ChainID, intent.Src))
		}

		if len(intent.Dst) != 0 && c.Insert(intent.ChainID, intent.Dst) {
			accounts = append(accounts, model.NewSessionAccount(intent.ChainID, intent.Dst))
		}
	}
	return
}

func (svc *Service) initISCAddress(
	intents []*opintent.TransactionIntent, accounts []*model.SessionAccount) (
	iscAddress []byte, err error) {
	var (
		txs       = make([][]byte, len(intents))
		owners    = make([][]byte, 0, len(accounts))
		signature uiptypes.Signature
	)
	for i, intent := range intents {
		txs[i] = intent.Bytes()
	}
	for _, owner := range accounts {
		owners = append(owners, owner.GetAddress())
	}
	if signature, err = svc.signer.Sign(bytes.Join(txs, []byte{})); err != nil {
		err = wrapper.Wrap(types.CodeSessionSignTxsError, err)
		return
	}
	if iscAddress, err = svc.nsbClient.CreateISC(svc.signer, make([]uint32, len(owners)), owners, txs, signature.Bytes()); err != nil {
		err = wrapper.Wrap(types.CodeSessionRequestNSBError, err)
		return
	}
	return
}
