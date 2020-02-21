package model

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	mcore "github.com/Myriad-Dreamin/go-ves/lib/backend/core"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

var p = NewDBLayerModule()

func GetInstance() *gorm.DB {
	return p.GetGormInstance()
}


func GetRawInstance() *sql.DB {
	return p.GetRawSQLInstance()
}

func GetDormInstance() *dorm.DB {
	return p.GetDormInstance()
}

func InstallFromContext(dep module.Module) bool {
	return p.FromContext(dep)
}

func Install(dep module.Module) bool {
	return p.Install(dep)
}

func InstallMock(dep module.Module, callback mcore.MockCallback) bool {
	return p.InstallMock(dep, callback)
}

func Close(dep module.Module) bool {
	return p.Close(dep)
}

func Configuration(cfg *config.ServerConfig) {
	p.Configuration(cfg)
}

func (p *DBLayerModule) Configuration(cfg *config.ServerConfig) {
	p.GetRawSQLInstance().SetMaxIdleConns(cfg.DatabaseConfig.MaxIdle)
	p.GetRawSQLInstance().SetMaxOpenConns(cfg.DatabaseConfig.MaxActive)
}

type VESDB interface {
	SetIndex(types.Index) (successOrNot bool)
	SetMultiIndex(types.MultiIndex) (successOrNot bool)
	SetSessionBase(types.SessionBase) (successOrNot bool)
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

	abstraction.UserBase

	SetKV(iscAddress []byte, provedKey []byte, provedValue []byte) error
	GetKV(iscAddress []byte, provedKey []byte) (provedValue []byte, err error)

	types.StorageHandlerInterface
	types.ChainDNSInterface
}
