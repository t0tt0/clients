package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/Myriad-Dreamin/go-ves/config"
	"github.com/Myriad-Dreamin/go-ves/lib/bni/getter"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io/ioutil"
	"os"
)

type OpIntents struct {
	Intents      []json.RawMessage `json:"op-intents"`
	Dependencies []json.RawMessage `json:"dependencies"`
}

func (o OpIntents) GetContents() (bs [][]byte) {
	xs := o.Intents
	for i := range xs {
		// fmt.Println(string(xs[i]))
		bs = append(bs, xs[i])
	}
	return
}

func (o OpIntents) GetDependencies() (bs [][]byte) {
	xs := o.Dependencies
	for i := range xs {
		// fmt.Println(string(xs[i]))
		bs = append(bs, xs[i])
	}
	return
}

func main() {
	sugar.WithFile(func(f *os.File) {
		var intents OpIntents
		sugar.HandlerError0(json.Unmarshal(sugar.HandlerError(ioutil.ReadAll(f)).([]byte), &intents))
		var ier = sugar.HandlerError(opintent.NewInitializer(config.UserMap, getter.NewBlockChainGetter(config.ChainDNS))).(*opintent.Initializer)
		res := sugar.HandlerError(ier.Parse(intents)).(opintent.TxIntents).GetTxIntents()
		for _, intent := range res {
			intent := intent.GetIntent()
			fmt.Println("=================================================================")
			fmt.Println("src:", hex.EncodeToString(intent.Src))
			fmt.Println("dst:", hex.EncodeToString(intent.Dst))
			fmt.Println("meta:", string(intent.Meta))
			fmt.Println("trans_type:", intent.TransType)
			fmt.Println("chain_id:", intent.ChainID)
			fmt.Println("amt:", intent.Amt)
		}
	}, "intent.json")
}
