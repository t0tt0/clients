package ethclient

import (
	"encoding/hex"
	"fmt"
	"testing"
)

const (
	testHost = "121.89.200.234:8545"
)

func TestGetEthAccounts(t *testing.T) {
	x, err := NewEthClient(testHost).GetEthAccounts()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(x)
}

func TestUnlock(t *testing.T) {
	ok, err := NewEthClient(testHost).PersonalUnlockAccout("0x6d8c6cb9d26b5a21ae498a22385ae4265f494cfc", "123456", 600)

	if ok == false || err != nil {
		if ok == false {
			if err != nil {

				t.Error(err)
			} else {
				t.Errorf("not ok..")
			}
		} else {

			t.Error(err)
		}
		return
	}
}

const objjj = `{"from":"0x6d8c6cb9d26b5a21ae498a22385ae4265f494cfc", "to": "0x6d8c6cb9d26b5a21ae498a22385ae4265f494cfc", "value": "0x1"}`

func TestSendTransaction(t *testing.T) {
	b, err := NewEthClient(testHost).SendTransaction([]byte(objjj))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestGetStorageAt(t *testing.T) {
	var addr = "1234567812345678123456781234567812345678"
	baddr, err := hex.DecodeString(addr)
	if err != nil {
		t.Error(err)
		return
	}
	var pos = []byte{1}
	b, err := NewEthClient(testHost).GetStorageAt(baddr, pos, "latest")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestGetTransactionByHash(t *testing.T) {
	txb, err := hex.DecodeString("a41d03fde4e7cf4c58870092c65709db7532956f7d0882156f11f503a6d88d2f")
	if err != nil {
		t.Error(err)
		return
	}
	b, err := NewEthClient(testHost).GetTransactionByHash(txb)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
	b, err = NewEthClient(testHost).GetTransactionByStringHash("0xa41d03fde4e7cf4c58870092c65709db7532956f7d0882156f11f503a6d88d2f")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))

}

func TestGetBlockByHash(t *testing.T) {
	txb, err := hex.DecodeString("8a8b9aaa48e0fb024abb7105798ad48057cf4fd14100505addabc319ed3d41c6")
	if err != nil {
		t.Error(err)
		return
	}

	b, err := NewEthClient(testHost).GetBlockByHash(txb, true)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))

	b, err = NewEthClient(testHost).GetBlockByHash(txb, false)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(b))
}
