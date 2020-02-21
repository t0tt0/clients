package bni

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/uip"
	nsbcli "github.com/HyperService-Consortium/go-ves/lib/net/nsb-client"
	"github.com/HyperService-Consortium/go-ves/lib/net/nsb-client/nsb-message"
	"net/url"
)

func (bn *BN) MustWithSigner() bool {
	return true
}

func (bn *BN) RouteWithSigner(signer uip.Signer) (uip.Router, error) {
	nbn := *bn
	nbn.signer = signer
	return &nbn, nil
}

func (bn *BN) RouteRaw(destination uip.ChainID, rawTransaction uip.RawTransaction) (
	transactionReceipt uip.TransactionReceipt, err error) {
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

func (bn *BN) WaitForTransact(_ uip.ChainID, transactionReceipt uip.TransactionReceipt,
	options ...interface{}) (blockID []byte, color []byte, err error) {
	var res nsb_message.ResultInfo
	err = json.Unmarshal(transactionReceipt, &res)
	if err != nil {
		return nil, nil, err
	}

	return []byte(res.Height), []byte(res.Hash), err
}

//func (bn *BN) Route(intent *uip.TransactionIntent, kvGetter uip.KVGetter) ([]byte, error) {
//	// todo
//	onChainTransaction, err := bn.Translate(intent, kvGetter)
//	if err != nil {
//		return nil, err
//	}
//	return bn.RouteRaw(intent.ChainID, onChainTransaction)
//}
