package bni

import (
	"bytes"
	"encoding/hex"
	"fmt"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/net/nsb-client"
	"github.com/gogo/protobuf/proto"
)

type rawTransaction struct {
	Type   transactiontype.Type
	Header *nsbrpc.TransactionHeader
}

func newRawTransaction(_type transactiontype.Type, header *nsbrpc.TransactionHeader) *rawTransaction {
	return &rawTransaction{Type: _type, Header: header}
}

func (r *rawTransaction) Serialize() ([]byte, error) {
	return nsbcli.GlobalClient.Serialize(r.Type, r.Header)
}

func (r *rawTransaction) Bytes() ([]byte, error) {
	b, err := proto.Marshal(r.Header)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *rawTransaction) Signed() bool {
	return len(r.Header.Signature) != 0
}

func (r *rawTransaction) Sign(user uip.Signer, ctxVars ...interface{}) (uip.RawTransaction, error) {
	if !bytes.Equal(r.Header.Src, user.GetPublicKey()) {
		return nil, fmt.Errorf("sign error user is %v, want is %v", hex.EncodeToString(user.GetPublicKey()), hex.EncodeToString(r.Header.Src))
	}
	r.Header = nsbcli.GlobalClient.Sign(user, r.Header)
	return r, nil
}
