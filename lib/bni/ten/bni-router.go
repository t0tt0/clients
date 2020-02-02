package bni

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"net/url"
)

func (bn *BN) MustWithSigner() bool {
	return true
}

func (bn *BN) RouteWithSigner(signer uiptypes.Signer) (uiptypes.Router, error) {
	nbn := *bn
	nbn.signer = signer
	return &nbn, nil
}

func (bn *BN) RouteRaw(destination uiptypes.ChainID, rawTransaction uiptypes.RawTransaction) (
	transactionReceipt uiptypes.TransactionReceipt, err error) {
	if !rawTransaction.Signed() {
		return nil, ErrNotSigned
	}
	ci, err := bn.dns.GetChainInfo(destination)
	if err != nil {
		return nil, err
	}
	// todo receipt
	b, err := rawTransaction.Bytes()
	if err != nil {
		return nil, err
	}
	b, err = nsbcli.NewNSBClient((&url.URL{Scheme: "http", Host: ci.GetChainHost(), Path: "/"}).String()).BroadcastTxCommitReturnBytes(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (bn *BN) WaitForTransact(_ uiptypes.ChainID, transactionReceipt uiptypes.TransactionReceipt,
	options ...interface{}) (blockID []byte, color []byte, err error) {
	var res nsbcli.ResultInfo
	err = json.Unmarshal(transactionReceipt, &res)
	if err != nil {
		return nil, nil, err
	}

	return []byte(res.Height), []byte(res.Hash), err
}

//func (bn *BN) Route(intent *uiptypes.TransactionIntent, kvGetter uiptypes.KVGetter) ([]byte, error) {
//	// todo
//	onChainTransaction, err := bn.Translate(intent, kvGetter)
//	if err != nil {
//		return nil, err
//	}
//	return bn.RouteRaw(intent.ChainID, onChainTransaction)
//}
