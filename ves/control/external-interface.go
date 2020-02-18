package control

import (
	"context"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	nsb_message "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client/nsb-message"
	"github.com/Myriad-Dreamin/go-ves/types"
	"google.golang.org/grpc"
)

type Signer = uiptypes.Signer
type ChainDNS interface {
	GetChainInfo(chainID uint64) (types.ChainInfo, error)
}

type CentralVESClient interface {
	InternalRequestComing(ctx context.Context, in *uiprpc.InternalRequestComingRequest, opts ...grpc.CallOption) (*uiprpc.InternalRequestComingReply, error)
	InternalAttestationSending(ctx context.Context, in *uiprpc.InternalRequestComingRequest, opts ...grpc.CallOption) (*uiprpc.InternalRequestComingReply, error)
	InternalCloseSession(ctx context.Context, in *uiprpc.InternalCloseSessionRequest, opts ...grpc.CallOption) (*uiprpc.InternalCloseSessionReply, error)
}
type Account = uiptypes.Account

type NSBClient interface {
	FreezeInfo(signer uiptypes.Signer, guid []byte, u uint64) ([]byte, error)
	AddMerkleProof(user uiptypes.Signer, toAddress []byte,
		merkleType uint16, rootHash, proof, key, value []byte) (*nsb_message.ResultInfo, error)
	AddBlockCheck(
		user uiptypes.Signer, toAddress []byte,
		chainID uint64, blockID, rootHash []byte, rcType uint8,
	) (*nsb_message.ResultInfo, error)
	InsuranceClaim(
		user uiptypes.Signer, contractAddress []byte,
		tid, aid uint64,
	) (*nsb_message.DeliverTx, error)
	CreateISC(signer uiptypes.Signer, uint32s []uint32, bytes [][]byte, bytes2 [][]byte, bytes3 []byte) ([]byte, error)
	SettleContract(signer uiptypes.Signer, bytes []byte) (*nsb_message.DeliverTx, error)
}
