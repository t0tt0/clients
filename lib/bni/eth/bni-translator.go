package bni

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (bn *BN) ParseTransactionIntent(intent uip.TxIntentI) (uip.TxIntentI, error) {
	return intent, nil
}

func (bn *BN) Translate(intent uip.TransactionIntent, storage uip.Storage) (uip.RawTransaction, error) {
	switch intent.GetTxType() {
	case trans_type.Payment:
		//meta := gjson.ParseBytes(intent.GetMeta())
		//value, err := payment_option.ParseInconsistentValueOption(meta, storage, intent.GetAmt())
		//if err != nil {
		//	return nil, err
		//}

		//fmt.Println(value, ".........")
//tyc: revise here......
		b, err := json.Marshal(map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "personal_sendTransaction",
			"params": []interface{}{
				map[string]interface{}{
					"from":  decoratePrefix(hex.EncodeToString(intent.GetSrc())),
					//"from":  decoratePrefix(hex.EncodeToString(intent.GetSrc())),
					"to":    decoratePrefix(hex.EncodeToString(intent.GetDst())),
					//"value": decorateValuePrefix(value),
					//06
					//"data": decorateValuePrefix(hex.EncodeToString([]byte("0x608060405233600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080819055506102fe8061005b6000396000f3fe6080604052600436106100345760003560e01c8063202ec1521461003957806382ab890a14610043578063da7d29821461006c575b600080fd5b610041610088565b005b34801561004f57600080fd5b5061006a6004803603810190610065919061022c565b61010f565b005b610086600480360381019061008191906101ff565b61016f565b005b6002600054141561010d57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc670de0b6b3a76400009081150290604051600060405180830381858888f19350505050158015610103573d6000803e3d6000fd5b5060066000819055505b565b6005600054148061012257506006600054145b8061012f57506001600054145b8061013c57506002600054145b1561014f5760005460008190555061016c565b600181148061015e5750600281145b1561016b57806000819055505b5b50565b600160005414156101d2578073ffffffffffffffffffffffffffffffffffffffff166108fc670de0b6b3a76400009081150290604051600060405180830381858888f193505050501580156101c8573d6000803e3d6000fd5b5060056000819055505b50565b6000813590506101e48161029a565b92915050565b6000813590506101f9816102b1565b92915050565b60006020828403121561021557610214610295565b5b6000610223848285016101d5565b91505092915050565b60006020828403121561024257610241610295565b5b6000610250848285016101ea565b91505092915050565b60006102648261026b565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b6102a381610259565b81146102ae57600080fd5b50565b6102ba8161028b565b81146102c557600080fd5b5056fea2646970667358221220dec7558ee50777622ebf86d566958cf3fa5d2016d23cf4949afebe49349bdc1464736f6c63430008070033"))),
				}, "1",
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
