package sessionservice

import (
	"encoding/json"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/ves/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func dataGoodTransactionIntent(t *testing.T) (
	*uiprpc_base.Account, *opintent.TransactionIntent, []byte) {
	t.Helper()
	var srcAcc = &uiprpc_base.Account{
		ChainId: 233,
		Address: []byte{2, 3, 3},
	}
	var ti = opintent.TransactionIntent{
		TransType: 0,
		Src:       srcAcc.Address,
		Dst:       nil,
		Meta:      nil,
		Amt:       "3e8",
		ChainID:   srcAcc.ChainId,
	}
	b, err := json.Marshal(&ti)
	if err != nil {
		t.Fatal("ser", err)
	}
	return srcAcc, &ti, b
}

func newMockGoodInternalPushTransaction(t *testing.T, f *fields, sesID []byte, sesFSet *mock.SessionFSet, cVes *mock.CentralVESClient) {
	srcAcc, _, b := dataGoodTransactionIntent(t)

	sesFSet.EXPECT().FindTransaction(sesID, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesID,
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(&uiprpc.InternalRequestComingReply{Ok:true}, nil)
}
