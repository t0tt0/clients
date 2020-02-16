package vesclient

import (
	mcore "github.com/Myriad-Dreamin/go-ves/lib/core"
	extend_traits "github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type traits = extend_traits.Traits

type traitsInterface interface {
	extend_traits.Interface
}

type traitsAcceptObject = extend_traits.ORMObject

func (m modelModule) newTraits(t traitsAcceptObject) traits {
	return extend_traits.NewTraits(t, m.GormDB, m.DormDB)
}

type modelModule struct {
	mcore.GormModule
	mcore.RawSQLModule
	mcore.DormModule
	mcore.LoggerModule

	accountTraits accountTraits
	sessionTraits sessionTraits
	Opened        bool
}

func newModelModule() modelModule {
	return modelModule{
		Opened: false,
	}
}

type DatabaseConfig struct {
	DataFilePath string
}

type DatabaseConfigGetter interface {
	GetVesClientDatabaseConfig() DatabaseConfig
}

func (m *modelModule) Install(dep module.Module) bool {
	dep.Provide(config.ModulePath.DBInstance.ModelModule, m)

	m.Opened = m.install(func(dep module.Module) bool {
		db, err := gorm.Open("sqlite3", dep.Require(config.ModulePath.Minimum.Global.Configuration).(DatabaseConfigGetter).GetVesClientDatabaseConfig().DataFilePath)
		if err != nil {
			m.Logger.Error("install sqlite error", "error", err)
			return false
		}
		return m.GormModule.FromRaw(db, dep)
	}, dep)
	return m.Opened
}

func (m *modelModule) install(
	initFunc func(dep module.Module) bool, dep module.Module) bool {
	return m.LoggerModule.Install(dep) &&
		initFunc(dep) &&
		m.RawSQLModule.FromRaw(m.GormDB.DB(), dep) &&
		m.DormModule.FromRawSQL(m.RawDB, dep) && mcore.ModelCallback(m, dep)
}
