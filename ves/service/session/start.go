package sessionservice

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/lib/uniquer"
	"github.com/HyperService-Consortium/go-ves/ves/model"
)

func (svc *Service) SessionStart(ctx context.Context, in *uiprpc.SessionStartRequest) (*uiprpc.SessionStartReply, error) {
	if instructions, accounts, err := svc.initOpIntents(in.GetOpintents()); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitOpIntentsError, err)
	} else {
		return svc.createSession(ctx, instructions, accounts)
	}
}

func (svc *Service) SessionStartR(ctx context.Context, in *uiprpc.SessionStartRequestR) (*uiprpc.SessionStartReply, error) {
	if instructions, accounts, err := svc.initOpIntentsR(in); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitOpIntentsError, err)
	} else {
		return svc.createSession(ctx, instructions, accounts)
	}
}

func (svc *Service) createSession(
	ctx context.Context,
	instructions []uip.Instruction, accounts []*model.SessionAccount) (*uiprpc.SessionStartReply, error) {
	var (
		ses        *model.Session
		iscAddress []byte
		err        error
	)

	if iscAddress, err = svc.initISCAddress(instructions, accounts); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitGUIDError, err)
	}
	if ses, err = svc.sesFSet.InitSessionInfo(iscAddress, instructions, accounts); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionInitError, err)
	}
	svc.logger.Info("new session requested", "address", hex.EncodeToString(ses.GetGUID()))

	// initializing accounts' bitmap in redis here a long time ago
	//if err = ses.InitAccountRedisMap(); err != nil {
	//	return nil, wrapper.Wrap(types.CodeSessionInitGUIDError, err)
	//}
	//

	for i := range instructions {
		//s.Logger.Info()
		if _, err = svc.nsbClient.FreezeInfo(svc.signer, ses.GetGUID(), uint64(i)); err != nil {
			return nil, wrapper.Wrap(types.CodeSessionFreezeInfoError, err)
		}
	}

	if err = svc.pushInternalInitRequestBySessionAccount(ctx, iscAddress, accounts); err != nil {
		return nil, err
	}

	return &uiprpc.SessionStartReply{
		Ok:        true,
		SessionId: iscAddress,
	}, nil
}

func (svc *Service) initOpIntents(opIntents opintent.OpIntents) (
	intents []uip.Instruction, accounts []*model.SessionAccount, err error) {
	var xIntents opintent.TxIntents
	xIntents, err = svc.opInitializer.Parse(opIntents)
	if err != nil {
		return
	}
	return svc.collectSessionInfomation(xIntents)
}

func (svc *Service) initOpIntentsR(opIntents opintent.OpIntentsPacket) (
	intents []uip.Instruction, accounts []*model.SessionAccount, err error) {
	var xIntents opintent.TxIntents
	xIntents, err = svc.opInitializer.ParseR(opIntents)
	if err != nil {
		return
	}
	return svc.collectSessionInfomation(xIntents)
}

func (svc *Service) collectSessionInfomation(xIntents opintent.TxIntents) (
	intents []uip.Instruction, accounts []*model.SessionAccount, err error) {

	var yIntents = xIntents.GetTxIntents()
	intents = make([]uip.Instruction, len(yIntents))
	for i := range yIntents {
		intents[i] = yIntents[i].GetInstruction()
	}
	c := uniquer.MakeUniquer()
	if c.Insert(svc.respAccount.GetChainId(), svc.respAccount.GetAddress()) {
		accounts = append(accounts, model.NewSessionAccount(svc.respAccount.GetChainId(), svc.respAccount.GetAddress()))
	}
	for _, intent := range intents {
		switch intent.GetType() {
		case instruction_type.Payment, instruction_type.ContractInvoke:

			//todo: remove assertion
			intent := intent.(*opintent.TransactionIntent)

			if c.Insert(intent.ChainID, intent.Src) {
				accounts = append(accounts, model.NewSessionAccount(intent.ChainID, intent.Src))
			}

			if len(intent.Dst) != 0 &&
				intent.TransType != trans_type.ContractInvoke && c.Insert(intent.ChainID, intent.Dst) {
				accounts = append(accounts, model.NewSessionAccount(intent.ChainID, intent.Dst))
			}
		}
	}
	return
}

func (svc *Service) initISCAddress(
	intents []uip.Instruction, accounts []*model.SessionAccount) (
	iscAddress []byte, err error) {
	var (
		owners    = make([][]byte, 0, len(accounts))
		signature uip.Signature
	)
	txs, err := opintent.EncodeInstructions(intents)
	if err != nil {
		return nil, err
	}
	for _, owner := range accounts {
		owners = append(owners, owner.GetAddress())
	}
	if signature, err = svc.signer.Sign(bytes.Join(txs, []byte{})); err != nil {
		err = wrapper.Wrap(types.CodeSessionSignTxsError, err)
		return
	}
	if iscAddress, err = svc.nsbClient.CreateISC(svc.signer, make([]uint64, len(owners)), owners, txs, signature.Bytes()); err != nil {
		err = wrapper.Wrap(types.CodeSessionRequestNSBError, err)
		return
	}
	return
}
