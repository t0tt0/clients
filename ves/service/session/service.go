//go:generate package-attach-to-path -generate_register_map
package sessionservice

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/Myriad-Dreamin/go-ves/ves/model/fset"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	cfg *config.ServerConfig
	key string

	accountDB      control.SessionAccountDBI
	db             control.SessionDBI
	sesFSet        control.SessionFSetI
	opInitializer  control.OpIntentInitializerI
	signer         control.Signer
	logger         types.Logger
	cVes           control.CentralVESClient
	respAccount    control.Account
	storageHandler control.StorageHandler
	dns            control.ChainDNS
	nsbClient      control.NSBClient

	// remove?
	storage control.SessionKV
}

func (svc *Service) UserRegister(context.Context, *uiprpc.UserRegisterRequest) (*uiprpc.UserRegisterReply, error) {
	panic("implement me")
}

func (svc *Service) SessionRequireTransact(context.Context, *uiprpc.SessionRequireTransactRequest) (*uiprpc.SessionRequireTransactReply, error) {
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

func (svc *Service) InformMerkleProof(context.Context, *uiprpc.MerkleProofReceiveRequest) (*uiprpc.MerkleProofReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) InformShortenMerkleProof(context.Context, *uiprpc.ShortenMerkleProofReceiveRequest) (*uiprpc.ShortenMerkleProofReceiveReply, error) {
	panic("implement me")
}

func (svc *Service) SessionServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.SessionService, error) {
	provider := m.Require(config.ModulePath.Minimum.Provider.Model).(*model.Provider)
	index := m.Require(config.ModulePath.DBInstance.Index).(types.Index)
	var a = &Service{
		key:       "sid",
		accountDB: provider.SessionAccountDB(),
		db:        provider.SessionDB(),
		sesFSet:   fset.NewSessionFSet(provider, index),

		dns:            m.Require(config.ModulePath.Service.ChainDNS).(control.ChainDNS),
		opInitializer:  m.Require(config.ModulePath.Service.OpIntentInitializer).(control.OpIntentInitializerI),
		cfg:            m.Require(config.ModulePath.Minimum.Global.Configuration).(*config.ServerConfig),
		logger:         m.Require(config.ModulePath.Minimum.Global.Logger).(types.Logger),
		cVes:           m.Require(config.ModulePath.Global.CentralVESClient).(control.CentralVESClient),
		signer:         m.Require(config.ModulePath.Global.Signer).(control.Signer),
		respAccount:    m.Require(config.ModulePath.Global.RespAccount).(control.Account),
		nsbClient:      m.Require(config.ModulePath.Global.NSBClient).(control.NSBClient),
		storageHandler: m.Require(config.ModulePath.Global.StorageHandler).(control.StorageHandler),
		storage:        m.Require(config.ModulePath.Global.Storage).(control.SessionKV),
	}
	return a, nil
}
