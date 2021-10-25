package bni

import (
	"errors"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	payment_option "github.com/HyperService-Consortium/go-ves/lib/bni/payment-option"
	"github.com/gogo/protobuf/proto"
	"github.com/tidwall/gjson"
)

func (bn *BN) ParseTransactionIntent(intent uip.TxIntentI) (uip.TxIntentI, error) {
	return intent, nil
}

func (bn *BN) Translate(x uip.TransactionIntent, storage uip.Storage) (uip.RawTransaction, error) {
	intent := x.(*opintent.TransactionIntent)

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
		var meta opintent.ContractInvokeMeta
		err := opintent.Serializer.Meta.Contract.Unmarshal(intent.GetMeta(), &meta)
		if err != nil {
			return nil, err
		}

		var faPair nsbrpc.FAPair
		faPair.FuncName = meta.FuncName
		//todo
		faPair.Args = nil
		header, err := nsbcli.GlobalClient.CreateUnsignedContractPacket(
			intent.Src, intent.Dst, []byte{0}, &faPair)
		if err != nil {
			return nil, err
		}
		return newRawTransaction(transactiontype.SendTransaction, header), nil
	default:
		return nil, errors.New("cant translate")
	}
}

func (bn *BN) Deserialize(raw []byte) (uip.RawTransaction, error) {

	var txHeader nsbrpc.TransactionHeader
	err := proto.Unmarshal(raw[1:], &txHeader)
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
