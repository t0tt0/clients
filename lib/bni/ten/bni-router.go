package bni

import (
	"encoding/json"
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	"github.com/HyperService-Consortium/go-uip/uip"
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

type Receipt struct {
	H []byte
	R []byte
}

func (bn *BN) RouteRaw(destination uip.ChainID, rawTransaction uip.RawTransaction) (
	transactionReceipt uip.TransactionReceipt, err error) {
	if !rawTransaction.Signed() {
		rawTransaction, err = rawTransaction.Sign(bn.signer)
		if err != nil {
			return nil, err
		}

		if !rawTransaction.Signed() {
			return nil, ErrNotSigned
		}
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
	v, err := nsbcli.NewNSBClient((&url.URL{Scheme: "http", Host: ci.GetChainHost(), Path: "/"}).String()).BroadcastTxCommitReturnBytes(b)
	if err != nil {
		return nil, err
	}

	b, err = json.Marshal(&Receipt{
		H: b[1:],
		R: v,
	})

	return b, err
}

func (bn *BN) WaitForTransact(_ uip.ChainID, transactionReceipt uip.TransactionReceipt,
	options ...interface{}) (blockID []byte, color []byte, err error) {
	var receipt Receipt
	err = json.Unmarshal(transactionReceipt, &receipt)
	if err != nil {
		return nil, nil, err
	}
	var res nsb_message.ResultInfo
	err = json.Unmarshal(receipt.R, &res)
	if err != nil {
		return nil, nil, err
	}

	return []byte(res.Height), receipt.H, err
}

//func (bn *BN) Route(intent *uip.TransactionIntent, kvGetter uip.KVGetter) ([]byte, error) {
//	// todo
//	onChainTransaction, err := bn.Translate(intent, kvGetter)
//	if err != nil {
//		return nil, err
//	}
//	return bn.RouteRaw(intent.ChainID, onChainTransaction)
//}
