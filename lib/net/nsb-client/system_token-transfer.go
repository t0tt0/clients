package nsbcli

import (
	"encoding/json"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-ves/lib/net/nsb-client/nsb-message"

	appl "github.com/HyperService-Consortium/NSB/application"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (nc *NSBClient) CreateTransferPacket(srcAddress, dstAddress []byte, value *math.Uint256) (*nsbrpc.TransactionHeader, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nc.transfer(value)
	if err != nil {
		return nil, err
	}
	txHeader, err := nc.CreateUnsignedContractPacket(srcAddress, dstAddress, value.Bytes(), fap)
	if err != nil {
		return nil, err
	}
	return txHeader, nil
}

func (nc *NSBClient) Transfer(
	user uip.Signer, toAddress []byte,
	value *math.Uint256,
) (*nsb_message.ResultInfo, error) {
	h, e := nc.CreateTransferPacket(user.GetPublicKey(), toAddress, value)
	return nc.systemCall(nc.sign(user, h, e))
}

func (nc *NSBClient) transfer(
	value *math.Uint256,
) (*nsbrpc.FAPair, error) {
	var args appl.ArgsTransfer
	args.Value = value
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "system.token@transfer"
	fap.Args = b
	// fmt.Println(PretiStruct(args), b)
	return fap, err
}
