package model

import (
	mcore "github.com/HyperService-Consortium/go-ves/lib/backend/core"
	extend_traits "github.com/HyperService-Consortium/go-ves/lib/backend/extend-traits"
	"github.com/HyperService-Consortium/go-ves/lib/basic/fcg"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	dblayer "github.com/HyperService-Consortium/go-ves/ves/model/internal/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type DBLayerModule struct {
	mcore.GormModule
	mcore.RawSQLModule
	mcore.DormModule
	mcore.LoggerModule

	Opened bool
}

func NewDBLayerModule() DBLayerModule {
	return DBLayerModule{
		Opened: false,
	}
}

func (p *DBLayerModule) FromContext(dep module.Module) bool {
	p.Opened = p.install(p.GormModule.FromContext, dep)
	return p.Opened
}

func (p *DBLayerModule) Install(dep module.Module) bool {
	p.Opened = p.install(p.GormModule.InstallFromConfiguration, dep)
	return p.Opened
}

func (p *DBLayerModule) InstallMock(dep module.Module, callback mcore.MockCallback) bool {
	p.Opened = p.install(p.GormModule.InstallMockFromConfiguration(callback), dep)
	if p.Opened {
		p.GormDB = p.GormDB.Debug()
	}
	return p.Opened
}

func (p *DBLayerModule) Close(dep module.Module) bool {
	if p.Opened {
		if err := p.GormDB.Close(); err != nil {
			p.Logger.Error("close DB error", "error", err)
			return false
		}
	}
	return true
}

func (p *DBLayerModule) NewTraits(t extend_traits.ORMObject) abstraction.ORMTraits {
	return p.newTraits(t)
}

func (p *DBLayerModule) newTraits(t interface{}) abstraction.ORMTraits {
	return extend_traits.NewTraits(t.(extend_traits.ORMObject), p.GormDB, p.DormDB)
}

func (p *DBLayerModule) Migrates(dep module.Module) error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//migrations
		p.migrate(dblayer.NewTransactionDB(p.newTraits, dep)),
		p.migrate(dblayer.NewSessionAccountDB(p.newTraits, dep)),
		p.migrate(dblayer.NewSessionDB(p.newTraits, dep)),
	})
}

func (p *DBLayerModule) install(
	initFunc func(dep module.Module) bool, dep module.Module) bool {
	return p.LoggerModule.Install(dep) &&
		initFunc(dep) &&
		p.RawSQLModule.FromRaw(p.GormDB.DB(), dep) &&
		p.DormModule.FromRawSQL(p.RawDB, dep) && mcore.Maybe(dep,
		"migrate callback error", p.Migrates(dep))
}

type canMigrate interface {
	Migrate() error
}

func (p *DBLayerModule) migrate(db canMigrate, err error) func() error {
	return func() error {
		if err != nil {
			return err
		} else {
			return db.Migrate()
		}
	}
}
