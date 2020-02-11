package vesclient

import (
	mcore "github.com/Myriad-Dreamin/go-ves/lib/core"
	extend_traits "github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

var p = newModelModule()

type traits = extend_traits.Traits

type traitsInterface interface {
	extend_traits.Interface
}

type traitsAcceptObject = extend_traits.ORMObject

func newTraits(t traitsAcceptObject) traits {
	return extend_traits.NewTraits(t, p.GormDB, p.DormDB)
}

type modelModule struct {
	mcore.GormModule
	mcore.RawSQLModule
	mcore.DormModule
	mcore.LoggerModule

	Opened bool
}

func newModelModule() modelModule {
	return modelModule{
		Opened: false,
	}
}

func (m *modelModule) Install(dep module.Module) bool {
	m.Opened = m.install(func(dep module.Module) bool {
		db, err := gorm.Open("sqlite3", "./test.db")
		if err != nil {
			m.Logger.Error("install sqlite error", "error",  err)
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
