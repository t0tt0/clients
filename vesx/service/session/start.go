package objectservice

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/types/session"
	"time"
)

//type MultiThreadSerialSessionStartService struct {
//	*vs.VServer
//	context.Context
//	*uiprpc.SessionStartRequest
//}
//
//func NewMultiThreadSerialSessionStartService(server *vs.VServer, context context.Context, sessionStartRequest *uiprpc.SessionStartRequest) MultiThreadSerialSessionStartService {
//	return MultiThreadSerialSessionStartService{VServer: server, Context: context, SessionStartRequest: sessionStartRequest}
//}
//
//func (s MultiThreadSerialSessionStartService) RequestNSBForNewSession(anyb types.Session) ([]byte, error) {
//	var accs = anyb.GetAccounts()
//
//	var owners = make([][]byte, 0, len(accs)+1)
//	// todo
//	// owners = append(owners, s.Signer.GetPublicKey())
//	for _, owner := range accs {
//		owners = append(owners, owner.GetAddress())
//		s.Logger.Info("waiting", hex.EncodeToString(owner.GetAddress()))
//	}
//	var txs = anyb.GetTransactions()
//	var btxs = make([][]byte, 0, len(txs))
//	for _, tx := range txs {
//		b, err := json.Marshal(tx)
//		if err != nil {
//			s.Logger.Error("error", "error", err)
//			return nil, err
//		}
//		btxs = append(btxs, b)
//	}
//
//	x, err := s.Signer.Sign(bytes.Join(anyb.GetTransactions(), []byte{}))
//	if err != nil {
//		return nil, err
//	}
//	return s.NsbClient.CreateISC(s.Signer, make([]uint32, len(owners)), owners, txs, x.Bytes())
//}
//
func (svc *sessionStartService) SessionStart() ([]byte, []uiptypes.Account, error) {
	var ses = session.MultiThreadSerialSession{Signer:svc.signer}
	success, helpInfo, err := ses.InitFromOpIntents(svc.GetOpintents())
	if err != nil {
		svc.logger.Error("error", "error", err)
		return nil, nil, err
	}
	if !success {
		return nil, nil, errors.New(helpInfo)
	}
	ses.ISCAddress, err = svc.RequestNSBForNewSession(ses)
	if ses.ISCAddress == nil {
		err = fmt.Errorf("request isc failed: %v", err)
		svc.Logger.Error("error", "error", err)
		return nil, nil, err
	}
	if err != nil {
		err = fmt.Errorf("request isc failed on request: %v", err)
		svc.Logger.Error("error", "error", err)
		return nil, nil, err
	}
	err = ses.AfterInitGUID()
	logger.Println("after init guid...", ses.ISCAddress, hex.EncodeToString(ses.ISCAddress))
	if err != nil {
		svc.Logger.Error("error", "error", err)
		return nil, nil, err
	}

	err = svc.DB.InsertSessionInfo(ses)
	if err != nil {
		svc.Logger.Error("error", "error", err)
		return nil, nil, err
	}
	for i := uint32(0); i < ses.TransactionCount; i++ {
		//s.Logger.Info()
		_, err := svc.NsbClient.FreezeInfo(svc.Signer, ses.ISCAddress, uint64(i))
		if err != nil {
			svc.Logger.Error("error", "error", err)
			return nil, nil, err
		}
	}

	// s.UpdateTxs
	// s.UpdateAccs
	return ses.ISCAddress, ses.GetAccounts(), nil
}

type sessionStartService struct {
	*Service
	context.Context
	*uiprpc.SessionStartRequest
}

func (svc *Service) Serve(ctx context.Context, in *uiprpc.SessionStartRequest) (*uiprpc.SessionStartReply, error) {
	s := sessionStartService{Service: svc, Context: ctx, SessionStartRequest: in}
	if b, accs, err := s.SessionStart(); err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		r, err := svc.CVes.InternalRequestComing(ctx, &uiprpc.InternalRequestComingRequest{
			SessionId: b,
			Host:      svc.Host,
			Accounts: func() (uaccs []*uipbase.Account) {
				for _, acc := range accs {
					uaccs = append(uaccs, &uipbase.Account{
						Address: acc.GetAddress(),
						ChainId: acc.GetChainId(),
					})
				}
				return
			}(),
		})

		if err != nil {
			svc.Logger.Error("error", "error", err)
			return nil, err
		}

		return &uiprpc.SessionStartReply{
			Ok:        r.GetOk(),
			SessionId: b,
		}, nil
	}
}



