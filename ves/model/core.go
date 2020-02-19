package model

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	mcore "github.com/Myriad-Dreamin/go-ves/lib/core"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
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
