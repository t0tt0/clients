package bni

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-ves/config"
	"testing"
)

func TestGetTransactionProof(t *testing.T) {
	txHeader, _ := hex.DecodeString("0a202333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff122076071c08f94940ad53f35bf1474610ae99569ec10e257be1426f19347d24667d1a01002220b34788fdbb73778ac4b7c69f6219cc06a59ab5dfb125ec4ab5ea928ef02fb57e2a060a04566f7465324062476a3a23ddee9c9da61255ef712693d87d83c4df4acd335cd3e23a81b6bc6d7e71523faa872283fd9b2c4526a354cfa3e89f78e9c055d9410ae212e102bc05")
	fmt.Println(NewBN(config.ChainDNS).GetTransactionProof(9, nil, txHeader))
}
