package vesclient

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/net/ves-websocket"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

type fields struct {
	p                      modelModule
	rwMutex                sync.RWMutex
	logger                 logger.Logger
	module                 DepModule
	closeSessionRWMutex    sync.RWMutex
	closeSessionSubscriber []SessionCloseSubscriber
	name                   []byte
	db                     AccountDBInterface
	conn                   ves_websocket.VESWSSocket
	nsbSigner              uiptypes.Signer
	dns                    types.ChainDNSInterface
	nsbClient              types.NSBClient
	waitOpt                uiptypes.RouteOptionTimeout
	cb                     chan *bytes.Buffer
	quit                   chan bool
	nsbip                  string
	grpcip                 string
	nsbBase                string
}

type fieldOptionDNS types.ChainDNSInterface
type fieldOptionNSBSigner uiptypes.Signer
type fieldOptionAccountDB AccountDBInterface
type fieldOptionNSBBase string

type fieldOption struct {
	dns       types.ChainDNSInterface
	nsbSigner uiptypes.Signer
	accountDB AccountDBInterface
	nsbBase   string
}

func withNSBBase(s string) fieldOptionNSBBase {
	return fieldOptionNSBBase(s)
}

func parseFieldOptions(rawOpts []interface{}) (options fieldOption) {
	options = fieldOption{
		dns: nil,
	}
	for i := range rawOpts {
		switch o := rawOpts[i].(type) {
		case fieldOptionDNS:
			options.dns = o
		case fieldOptionNSBSigner:
			options.nsbSigner = o
		case fieldOptionAccountDB:
			options.accountDB = o
		case fieldOptionNSBBase:
			options.nsbBase = string(o)
		}
	}
	return
}

var testLogger logger.Logger

func createFields(rawOpts ...interface{}) fields {
	options := parseFieldOptions(rawOpts)
	if testLogger == nil {
		var err error
		testLogger, err = logger.NewZapLogger(
			logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
		if err != nil {
			log.Fatal("init vesLogger error", "error", err)
		}
	}
	return fields{
		p:                      newModelModule(),
		rwMutex:                sync.RWMutex{},
		logger:                 testLogger,
		module:                 DepModule{},
		closeSessionRWMutex:    sync.RWMutex{},
		closeSessionSubscriber: nil,
		name:                   nil,
		db:                     options.accountDB,
		conn:                   nil,
		nsbSigner:              options.nsbSigner,
		dns:                    options.dns,
		nsbClient:              nil,
		waitOpt:                0,
		cb:                     nil,
		quit:                   nil,
		nsbip:                  "",
		grpcip:                 "",
		nsbBase:                options.nsbBase,
	}
}
