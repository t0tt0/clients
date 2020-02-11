package vesclient

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	core_cfg "github.com/Myriad-Dreamin/go-ves/lib/core-cfg"
	"github.com/Myriad-Dreamin/go-ves/lib/database/filedb"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
)

// VesClient is the web socket client interactive with veses
type VesClient struct {
	rwMutex sync.RWMutex
	logger  logger.Logger
	module  module.Module

	name   []byte
	signer uiptypes.Signer
	keys   *ECCKeys
	accs   *EthAccounts

	conn      *websocket.Conn
	nsbClient types.NSBClient
	waitOpt   uiptypes.RouteOptionTimeout

	cb   chan *bytes.Buffer
	quit chan bool

	fdb *filedb.FileDB

	nsbip  string
	grpcip string

	closeSessionRWMutex    sync.RWMutex
	closeSessionSubscriber []SessionCloseSubscriber
}

type cfgX struct{}
func (cfgX) GetDatabaseConfiguration() core_cfg.DatabaseConfig {
	return core_cfg.DatabaseConfig{
		Escaper:        `"`,
	}
}

var nilC cfgX

// NewVesClient return a pointer of VesClinet
func NewVesClient(rOptions ...interface{}) (vc *VesClient, err error) {
	options := parseOptions(rOptions)
	vc = &VesClient{
		cb:        make(chan *bytes.Buffer, 1),
		quit:      make(chan bool, 1),
		module:    make(module.Module),
		nsbClient: nsbcli.NewNSBClient(options.nsbHost),
		logger:    options.logger,
		waitOpt:   options.waitOpt,
		name:      options.vesName,
	}
	vc.module.Provide(config.ModulePath.Minimum.Global.Configuration, nilC)
	vc.module.Provide(config.ModulePath.Minimum.Global.Logger, vc.logger)
	if !p.Install(vc.module) {
		err = errInitModel
		return
	}

	vc.conn, _, err = new(websocket.Dialer).Dial((&url.URL{Scheme: "ws", Host: options.addr, Path: "/"}).String(), nil)
	return
}
