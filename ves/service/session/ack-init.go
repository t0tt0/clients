package sessionservice

import (
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"golang.org/x/net/context"
)

func (svc *Service) SessionAckForInit(ctx context.Context, in *uiprpc.SessionAckForInitRequest) (*uiprpc.SessionAckForInitReply, error) {
	//s.Logger.Info("session acknowledging... ", "address", hex.EncodeToString(s.GetUser().GetAddress()))
	ses, err := svc.getSession(in.GetSessionId())
	if err != nil {
		return nil, err
	}

	// todo: get Session Acked from isc
	if err = svc.sesFSet.AckForInit(
		ses, in.GetUser(),
		signaturer.FromRaw(in.GetUserSignature().Content,
			in.GetUserSignature().SignatureType)); err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAcknowledgeError, err)
	}
	c, err := svc.accountDB.GetAcknowledged(ses.ISCAddress)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAccountGetAcknowledgedError, err)
	}

	if c == ses.AccountsCount {

		pc, err := svc.nsbClient.ISCGetPC(svc.signer, ses.GetGUID())
		if err != nil {
			// todo: CodeGetPCError
			return nil, wrapper.Wrap(types.CodeSelectError, err)
		}

		//todo: remove conversion
		ses.UnderTransacting = int64(pc)

		if err = svc.pushTransaction(ctx, ses, ses.UnderTransacting); err != nil {
			return nil, err
		}
	}

	//if _, err = ses.Update(); err != nil {
	//	return nil, wrapper.Wrap(types.CodeUpdateError, err)
	//}

	return &uiprpc.SessionAckForInitReply{
		Ok: true,
	}, nil
}

// request -> start -> pushInternalInitRequestBySessionAccount -> acks -> ackForInit
