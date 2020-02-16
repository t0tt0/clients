package sessionservice

import (
	"encoding/json"
	"errors"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
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
	}).Return(&uiprpc.InternalRequestComingReply{Ok: true}, nil)
}

const (
	ethereumChainID uiptypes.ChainTypeUnderlyingType = iota
	tendermintChainID
	unknownChainTypeID
	unknownChainID
)

func newMockDNS(f *fields, dns *mock.ChainDNS) {
	dns.EXPECT().GetChainInfo(ethereumChainID).Return(ChainInfo{
		ChainType: ChainType.Ethereum,
		ChainHost: "orz.cc:23333",
	}, nil)
	dns.EXPECT().GetChainInfo(tendermintChainID).Return(ChainInfo{
		ChainType: ChainType.TendermintNSB,
		ChainHost: "orz.cc:23332",
	}, nil)
	dns.EXPECT().GetChainInfo(unknownChainTypeID).Return(ChainInfo{
		ChainType: ChainType.Unassigned,
	}, nil).MinTimes(0)
	dns.EXPECT().GetChainInfo(unknownChainID).Return(nil, errors.New("not found"))
}
