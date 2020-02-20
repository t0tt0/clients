package bni

import (
	"errors"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	payment_option "github.com/Myriad-Dreamin/go-ves/lib/bni/payment-option"
	"github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"github.com/gogo/protobuf/proto"
	"github.com/tidwall/gjson"
)

func (bn *BN) ParseTransactionIntent(intent uip.TxIntentI) (uip.TxIntentI, error) {
	panic("implement me")
}

func (bn *BN) Translate(intent *uip.TransactionIntent, storage uip.Storage) (uip.RawTransaction, error) {
	switch intent.TransType {
	case trans_type.Payment:
		meta := gjson.ParseBytes(intent.Meta)
		value, err := payment_option.ParseInconsistentValueOption(meta, storage, intent.Amt)
		if err != nil {
			return nil, err
		}
		header, err := nsbcli.GlobalClient.CreateTransferPacket(intent.Src, intent.Dst, math.NewUint256FromHexString(value))
		if err != nil {
			return nil, err
		}
		return newRawTransaction(transactiontype.SystemCall, header), nil
	case trans_type.ContractInvoke:
		// var meta uip.ContractInvokeMeta
		//
		// err := json.Unmarshal(intent.Meta, &meta)
		// if err != nil {
		// 	return nil, err
		// }
		// //_ = meta
		//
		// data, err := ContractInvocationDataABI(&meta, storage)
		// if err != nil {
		// 	return nil, err
		// }
		//
		// hexdata := hex.EncodeToString(data)
		// // meta.FuncName
		//
		// return json.Marshal(map[string]interface{}{
		// 	"jsonrpc": "2.0",
		// 	"method":  "eth_sendTransaction",
		// 	"params": []interface{}{
		// 		map[string]interface{}{
		// 			"from":  decoratePrefix(hex.EncodeToString(intent.Src)),
		// 			"to":    decoratePrefix(hex.EncodeToString(intent.Dst)),
		// 			"value": decoratePrefix(intent.Amt),
		// 			"data":  decoratePrefix(hexdata),
		// 		},
		// 	},
		// 	"id": 1,
		// })
		return nil, errors.New("todo")
	default:
		return nil, errors.New("cant translate")
	}
}

func (bn *BN) Deserialize(raw []byte) (uip.RawTransaction, error) {

	var txHeader nsbrpc.TransactionHeader
	err := proto.Unmarshal(raw, &txHeader)
	if err != nil {
		return nil, err
	}
	if len(txHeader.Src) != 32 {
		return nil, ErrorDecodeSrcAddress
	}
	if len(txHeader.Dst) != 32 && len(txHeader.Dst) != 0 {
		return nil, ErrorDecodeDstAddress
	}

	return &rawTransaction{
		Type:   raw[0],
		Header: &txHeader,
	}, nil
}
