package sessionservice

import (
	"context"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/types/nsb-interface"
)

type AttestationAdapdator struct {
	*uiprpc_base.Attestation
}

func (atte *AttestationAdapdator) GetSignatures() []uip.Signature {
	var ss = atte.Attestation.GetSignatures()
	ret := make([]uip.Signature, len(ss))
	for _, s := range ss {
		ret = append(ret, signaturer.FromRaw(s.Content, s.SignatureType))
	}
	return ret
}

func (svc *Service) InformAttestation(ctx context.Context, in *uiprpc.AttestationReceiveRequest) (*uiprpc.AttestationReceiveReply, error) {
	//svc.logger.Info("session init acknowledging... ", "address", hex.EncodeToString(svc.accountDB.ID(ses.ISCAddress))
	//svc.logger.Info("into last function")
	ses, err := svc.getSession(in.GetSessionId())

	//svc.logger.Info("session inform attestation", "address", )

	if err != nil {
		return nil, err
	}

	_, bn, err := svc.getCurrentTxIntentWithExecutor(ses)
	if err != nil {
		return nil, err
	}
	lastTxID := ses.UnderTransacting
	err = svc.sesFSet.NotifyAttestation(
		ses,
		nsbi.NSBInterfaceFromClient(svc.nsbClient, svc.signer),
		bn,
		&AttestationAdapdator{Attestation: in.GetAtte()},
	)



	if err != nil {
		svc.logger.Info("i think  may be here.................")
		return nil, wrapper.Wrap(types.CodeSessionNotifyAttestationError, err)
	}


	//svc.logger.Info("before get", "undertransacting, is for informattestation", ses.UnderTransacting)
	//svc.logger.Info("before get", "transactioncount, is for informattestation", ses.TransactionCount)

	if ses.UnderTransacting >= ses.TransactionCount {
		// close
		svc.logger.Info("intohere.........................?")
		accounts, err := svc.accountDB.ID(ses.ISCAddress)
		if err != nil {
			return nil, wrapper.Wrap(types.CodeSessionAccountGetTotalError, err)
		}

		if err = svc.pushInternalCloseRequestBySessionAccount(ctx, in.GetSessionId(), accounts); err != nil {
			return nil, err
		}
	}

	//svc.logger.Info("ses.undertransacting transid", "transid is for informattestation", ses.UnderTransacting)
	//svc.logger.Info("lastTxID", "lasttransid is for informattestation", lastTxID)


	if ses.UnderTransacting != lastTxID {
		//svc.logger.Info("just test whether logger is added............")
		//svc.logger.Info("will go into relay & push")
		if err = svc.pushTransaction(ctx, ses, ses.UnderTransacting); err != nil {
			return nil, err
		}
	}
	return &uiprpc.AttestationReceiveReply{
		Ok: true,
	}, nil
}
