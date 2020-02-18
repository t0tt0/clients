package dblayer

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/lib/core"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/go-ves/lib/fcg"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var p = newModelModule()

func GetInstance() *gorm.DB {
	return p.GetGormInstance()
}

func GetRawInstance() *sql.DB {
	return p.GetRawSQLInstance()
}

func GetDormInstance() *dorm.DB {
	return p.GetDormInstance()
}

func FromContext(dep module.Module) bool {
	return p.FromContext(dep)
}

func Install(dep module.Module) bool {
	return p.Install(dep)
}

func InstallMock(dep module.Module, callback mcore.MockCallback) bool {
	return p.InstallMock(dep, callback)
}

func Close(dep module.Module) bool {
	if p.Opened {
		if err := p.GormDB.Close(); err != nil {
			p.Logger.Error("close DB error", "error", err)
			return false
		}
	}
	return true
}

func Configuration(cfg *config.ServerConfig) {
	(*p.RawDB).SetMaxIdleConns(cfg.DatabaseConfig.MaxIdle)
	(*p.RawDB).SetMaxOpenConns(cfg.DatabaseConfig.MaxActive)
}

type Traits = extend_traits.Traits

type Interface interface {
	extend_traits.Interface
}

type TraitsAcceptObject = extend_traits.ORMObject

func NewTraits(t TraitsAcceptObject) Traits {
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

func (m *modelModule) install(
	initFunc func(dep module.Module) bool, dep module.Module) bool {
	return m.LoggerModule.Install(dep) &&
		initFunc(dep) &&
		m.RawSQLModule.FromRaw(m.GormDB.DB(), dep) &&
		m.DormModule.FromRawSQL(m.RawDB, dep) && mcore.ModelCallback(m, dep)
}

func (m *modelModule) FromContext(dep module.Module) bool {
	m.Opened = m.install(m.GormModule.FromContext, dep)
	return m.Opened
}

func (m *modelModule) Install(dep module.Module) bool {
	m.Opened = m.install(m.GormModule.InstallFromConfiguration, dep)
	if m.Opened {
		m.GormDB = m.GormDB.Debug()
	}
	return m.Opened
}

func (m *modelModule) InstallMock(dep module.Module, callback mcore.MockCallback) bool {
	m.Opened = m.install(m.GormModule.InstallMockFromConfiguration(callback), dep)
	if m.Opened {
		m.GormDB = m.GormDB.Debug()
	}
	return m.Opened
}

func (modelModule) Migrates() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//migrations
		SessionAccount{}.migrate,
		Session{}.migrate,
	})
}

func (modelModule) Injects() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//injections
		injectSessionAccountTraits,
		injectSessionTraits,
	})
}

func decodeBase64(src string) []byte {
	b, err := encoding.DecodeBase64(src)
	if err != nil {
		p.Logger.Debug("decode failed", "error", err, "source", src)
		return nil
	}
	return b
}

func DecodeAddress(src string) []byte {
	return decodeBase64(src)
}

func EncodeAddress(src []byte) string {
	return encoding.EncodeBase64(src)
}
