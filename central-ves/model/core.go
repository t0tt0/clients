package model

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/db-layer"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/sp-layer"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

func InstallFromContext(dep module.Module) bool {
	return dblayer.FromContext(dep)
}

func Install(dep module.Module) bool {
	return dblayer.Install(dep)
}

func InstallMock(dep module.Module) bool {
	return dblayer.InstallMock(dep)
}

func RegisterRedis(dep module.Module) bool {
	return splayer.RegisterRedis(dep)
}

func InstallRedis(dep module.Module) bool {
	return splayer.Install(dep)
}

func Close(dep module.Module) bool {
	x := dblayer.Close(dep)
	x = x && splayer.Close(dep)
	return x
}

type Provider = splayer.Provider

func NewProvider(namespace string) *Provider {
	return splayer.NewProvider(namespace)
}

func SetProvider(p *Provider) *Provider {
	return splayer.SetProvider(p)
}

type VESDB interface {
	SetIndex(types.Index) (successOrNot bool)
	SetMultiIndex(types.MultiIndex) (successOrNot bool)
	SetSessionBase(types.SessionBase) (successOrNot bool)
	SetSessionKVBase(types.SessionKVBase) (successOrNot bool)
	SetStorageHandler(types.StorageHandler) (successOrNot bool)
	SetChainDNS(types.ChainDNS) (successOrNot bool)

	// insert accounts maps from guid to account
	InsertSessionInfo(types.Session) error

	// find accounts which guid is corresponding to user
	FindSessionInfo(iscAddress []byte) (types.Session, error)

	UpdateSessionInfo(types.Session) error

	DeleteSessionInfo(iscAddress []byte) error

	FindTransaction(iscAddress []byte, transactionID uint64, callback func([]byte) error) error

	ActivateSession(iscAddress []byte)

	InactivateSession(iscAddress []byte)

	UserBase

	SetKV(iscAddress []byte, provedKey []byte, provedValue []byte) error
	GetKV(iscAddress []byte, provedKey []byte) (provedValue []byte, err error)

	GetSetter(iscAddress []byte) types.KVSetter
	GetGetter(iscAddress []byte) types.KVGetter

	types.StorageHandlerInterface
	types.ChainDNSInterface
}
