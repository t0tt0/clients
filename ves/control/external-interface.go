package control

import (
	"context"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	nsb_message "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client/nsb-message"
	"github.com/Myriad-Dreamin/go-ves/types"
	"google.golang.org/grpc"
)

type Signer = uip.Signer
type ChainDNS interface {
	GetChainInfo(chainID uint64) (types.ChainInfo, error)
}

type CentralVESClient interface {
	InternalRequestComing(ctx context.Context, in *uiprpc.InternalRequestComingRequest, opts ...grpc.CallOption) (*uiprpc.InternalRequestComingReply, error)
	InternalAttestationSending(ctx context.Context, in *uiprpc.InternalRequestComingRequest, opts ...grpc.CallOption) (*uiprpc.InternalRequestComingReply, error)
	InternalCloseSession(ctx context.Context, in *uiprpc.InternalCloseSessionRequest, opts ...grpc.CallOption) (*uiprpc.InternalCloseSessionReply, error)
}
type Account = uip.Account

type NSBClient interface {
	FreezeInfo(signer uip.Signer, guid []byte, u uint64) ([]byte, error)
	AddMerkleProof(user uip.Signer, toAddress []byte,
		merkleType uint16, rootHash, proof, key, value []byte) (*nsb_message.ResultInfo, error)
	AddBlockCheck(
		user uip.Signer, toAddress []byte,
		chainID uint64, blockID, rootHash []byte, rcType uint8,
	) (*nsb_message.ResultInfo, error)
	InsuranceClaim(
		user uip.Signer, contractAddress []byte,
		tid, aid uint64,
	) (*nsb_message.DeliverTx, error)
	CreateISC(signer uip.Signer, uint32s []uint32, bytes [][]byte, bytes2 [][]byte, bytes3 []byte) ([]byte, error)
	SettleContract(signer uip.Signer, bytes []byte) (*nsb_message.DeliverTx, error)
}
