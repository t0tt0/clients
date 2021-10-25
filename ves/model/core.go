package model

import (
	"database/sql"
	mcore "github.com/HyperService-Consortium/go-ves/lib/backend/core"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/dorm"
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

type ORMTraits = abstraction.ORMTraits
type SessionAccountQueryOption = abstraction.SessionAccountQueryOption
type SessionQueryOption = abstraction.SessionQueryOption
