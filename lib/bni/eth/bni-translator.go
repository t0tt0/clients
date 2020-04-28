package bni

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	payment_option "github.com/HyperService-Consortium/go-ves/lib/bni/payment-option"
	"github.com/tidwall/gjson"
)

func (bn *BN) ParseTransactionIntent(intent uip.TxIntentI) (uip.TxIntentI, error) {
	return intent, nil
}

func (bn *BN) Translate(intent uip.TransactionIntent, storage uip.Storage) (uip.RawTransaction, error) {
	switch intent.GetTxType() {
	case trans_type.Payment:
		meta := gjson.ParseBytes(intent.GetMeta())
		value, err := payment_option.ParseInconsistentValueOption(meta, storage, intent.GetAmt())
		if err != nil {
			return nil, err
		}

		//fmt.Println(value, ".........")

		b, err := json.Marshal(map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_sendTransaction",
			"params": []interface{}{
				map[string]interface{}{
					"from":  decoratePrefix(hex.EncodeToString(intent.GetSrc())),
					"to":    decoratePrefix(hex.EncodeToString(intent.GetDst())),
					"value": decorateValuePrefix(value),
				},
			},
			"id": 1,
		})
		//fmt.Println("...", string(b))
		return NewRawTransaction(b, intent.GetSrc(), false), err
	case trans_type.ContractInvoke:
		var meta opintent.ContractInvokeMeta
		err := opintent.Serializer.Meta.Contract.Unmarshal(intent.GetMeta(), &meta)
		if err != nil {
			return nil, err
		}

		data, err := ContractInvocationDataABI(intent.GetChainID(), &meta, storage)
		if err != nil {
			return nil, err
		}

		hexdata := hex.EncodeToString(data)

		b, err := json.Marshal(map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_sendTransaction",
			"params": []interface{}{
				map[string]interface{}{
					"from": decoratePrefix(hex.EncodeToString(intent.GetSrc())),
					"to":   decoratePrefix(hex.EncodeToString(intent.GetDst())),
					// todo
					"gas": "0x7a1200",
					// todo
					//"value": decoratePrefix(intent.Amt),
					"data": decorateValuePrefix(hexdata),
				},
			},
			"id": 1,
		})
		return NewRawTransaction(b, intent.GetSrc(), false), err
	default:
		return nil, errors.New("cant translate")
	}
}

func (bn *BN) Deserialize(raw []byte) (rawTransaction uip.RawTransaction, err error) {
	var x = new(RawTransaction)
	err = json.Unmarshal(raw, x)
	return x, err
}
