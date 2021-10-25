package vesclient

import (
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

type DepModule struct {
	module.Module
}

type Module = DepModule

func newDepModule() DepModule {
	return DepModule{Module: make(module.Module)}
}

func (dep DepModule) GormDB() *gorm.DB {
	return dep.Module.Require(config.ModulePath.Minimum.DBInstance.GormDB).(*gorm.DB)
}

func (dep DepModule) ModelModule() *modelModule {
	return dep.Module.Require(config.ModulePath.DBInstance.ModelModule).(*modelModule)
}
