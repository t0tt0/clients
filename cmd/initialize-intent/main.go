package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/config"
	"github.com/HyperService-Consortium/go-ves/lib/bni/getter"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io/ioutil"
	"os"
)

var (
	fp = flag.String("i", "intent.json", "intent file path")
)

type packet []byte

func (p packet) GetContent() (content []byte) {
	return p
}

func desc(code instruction_type.Type) string {
	switch code {
	case instruction_type.ContractInvoke:
		return "ContractInvoke"
	case instruction_type.Payment:
		return "Payment"
	default:
		return fmt.Sprintf("InstType(%d)", code)
	}
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	sugar.WithFile(func(f *os.File) {
		var ier = sugar.HandlerError(opintent.NewInitializer(config.UserMap, getter.NewBlockChainGetter(config.ChainDNS))).(*opintent.Initializer)
		res := sugar.HandlerError(ier.ParseR(packet(
			sugar.HandlerError(ioutil.ReadAll(f)).([]byte)))).(opintent.TxIntents).GetTxIntents()
		for _, intent := range res {
			intent := intent.GetInstruction()
			fmt.Println("=================================================================")

			fmt.Println("instruction_type:", desc(intent.GetType()))
			if intent.GetType() <= trans_type.ContractInvoke {
				intent := intent.(uip.TransactionIntent)
				fmt.Println("src:", hex.EncodeToString(intent.GetSrc()))
				fmt.Println("dst:", hex.EncodeToString(intent.GetDst()))
				fmt.Println("meta:", string(intent.GetMeta()))
				fmt.Println("trans_type:", desc(intent.GetTxType()))
				fmt.Println("chain_id:", intent.GetChainID())
				fmt.Println("amt:", intent.GetAmt())
			}
		}
	}, *fp)
}
