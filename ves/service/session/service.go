//go:generate package-attach-to-path -generate_register_map
package sessionservice

import (
	"context"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/Myriad-Dreamin/go-ves/ves/model/fset"
	"github.com/Myriad-Dreamin/go-ves/ves/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	db             *model.SessionDB
	sesFSet        *fset.SessionFSet
	signer         uiptypes.Signer
	cfg            *config.ServerConfig
	logger         types.Logger
	key            string
	cVes           uiprpc.CenteredVESClient
	nsbClient      types.NSBClient
	accountDB      *splayer.SessionAccountDB
	opInitializer  *opintent.OpIntentInitializer
	respAccount    uiptypes.Account
	storage        types.SessionKV
	storageHandler types.StorageHandler
	dns types.ChainDNSInterface
}

func (svc *Service) UserRegister(context.Context, *uiprpc.UserRegisterRequest) (*uiprpc.UserRegisterReply, error) {
	panic("implement me")
}

func (svc *Service) SessionRequireTransact(context.Context, *uiprpc.SessionRequireTransactRequest) (*uiprpc.SessionRequireTransactReply, error) {
	panic("implement me")
}

func (svc *Service) SessionRequireRawTransact(context.Context, *uiprpc.SessionRequireRawTransactRequest) (*uiprpc.SessionRequireRawTransactReply, error) {
	panic("implement me")
}

func (svc *Service) AttestationReceive(context.Context, *uiprpc.AttestationReceiveRequest) (*uiprpc.AttestationReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) MerkleProofReceive(context.Context, *uiprpc.MerkleProofReceiveRequest) (*uiprpc.MerkleProofReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) ShrotenMerkleProofReceive(context.Context, *uiprpc.ShortenMerkleProofReceiveRequest) (*uiprpc.ShortenMerkleProofReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) InformAttestation(context.Context, *uiprpc.AttestationReceiveRequest) (*uiprpc.AttestationReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) InformMerkleProof(context.Context, *uiprpc.MerkleProofReceiveRequest) (*uiprpc.MerkleProofReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) InformShortenMerkleProof(context.Context, *uiprpc.ShortenMerkleProofReceiveRequest) (*uiprpc.ShortenMerkleProofReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) SessionServiceSignatureXXX() interface{} { return svc }


func NewService(m module.Module) (control.SessionService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Minimum.Provider.Model).(*model.Provider)
	a.accountDB = provider.SessionAccountDB()
	a.dns = m.Require(config.ModulePath.Service.ChainDNS).(types.ChainDNSInterface)
	a.cVes = m.Require(config.ModulePath.Global.CentralVESClient).(uiprpc.CenteredVESClient)
	a.cfg = m.Require(config.ModulePath.Minimum.Global.Configuration).(*config.ServerConfig)
	a.logger = m.Require(config.ModulePath.Minimum.Global.Logger).(types.Logger)
	a.signer = m.Require(config.ModulePath.Global.Signer).(uiptypes.Signer)
	a.respAccount = m.Require(config.ModulePath.Global.RespAccount).(uiptypes.Account)
	index := m.Require(config.ModulePath.DBInstance.Index).(types.Index)
	a.nsbClient = m.Require(config.ModulePath.Global.NSBClient).(types.NSBClient)
	a.storageHandler = m.Require(config.ModulePath.Global.StorageHandler).(types.StorageHandler)
	a.storage = m.Require(config.ModulePath.Global.Storage).(types.SessionKV)
	a.opInitializer = m.Require(config.ModulePath.Service.OpIntentInitializer).(*opintent.OpIntentInitializer)
	a.sesFSet = fset.NewSessionFSet(provider, index)
	a.key = "sid"
	a.db = provider.SessionDB()
	return a, nil
}
