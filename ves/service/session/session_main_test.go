package sessionservice

import (
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/mock"
	"github.com/golang/mock/gomock"
)

type fields struct {
	cfg            *config.ServerConfig
	key            string
	accountDB      control.SessionAccountDBI
	db             control.SessionDBI
	sesFSet        control.SessionFSetI
	opInitializer  control.OpIntentInitializerI
	signer         control.Signer
	logger         control.Logger
	cVes           control.CenteredVESClient
	respAccount    control.Account
	storage        control.SessionKV
	storageHandler control.StorageHandler
	dns            control.ChainDNSInterface
	nsbClient      control.NSBClient
}

func MockSessionDB(ctl *gomock.Controller) *mock.SessionDB {
	return mock.NewSessionDB(ctl)
}

func MockSessionFSet(ctl *gomock.Controller) *mock.SessionFSet {
	return mock.NewSessionFSet(ctl)
}

func MockSessionAccountDB(ctl *gomock.Controller) *mock.SessionAccountDB {
	return mock.NewSessionAccountDB(ctl)
}

func createField(options ...interface{}) fields {
	f := fields{}
	for i := range options {
		switch o := options[i].(type) {
		case *mock.SessionDB:
			f.db = o
		case *mock.SessionFSet:
			f.sesFSet = o
		case *mock.SessionAccountDB:
			f.accountDB = o
		}
	}
	return f
}



