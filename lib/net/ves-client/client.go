package vesclient

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	xconfig "github.com/Myriad-Dreamin/go-ves/config"
	core_cfg "github.com/Myriad-Dreamin/go-ves/lib/core-cfg"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
)

// VesClient is the web socket client interact with VESs
type VesClient struct {
	p                      modelModule
	rwMutex                sync.RWMutex
	logger                 logger.Logger
	module                 DepModule
	closeSessionRWMutex    sync.RWMutex
	closeSessionSubscriber []SessionCloseSubscriber

	name []byte

	db   AccountDBInterface
	conn SocketConn

	nsbSigner uiptypes.Signer
	dns       types.ChainDNSInterface
	nsbClient types.NSBClient

	waitOpt uiptypes.RouteOptionTimeout

	cb   chan *bytes.Buffer
	quit chan bool

	nsbip  string
	grpcip string

	nsbBase string
}

type cfgX struct {
	dc DatabaseConfig
}

func (c cfgX) GetVesClientDatabaseConfig() DatabaseConfig {
	return c.dc
}

func (cfgX) GetDatabaseConfiguration() core_cfg.DatabaseConfig {
	return core_cfg.DatabaseConfig{
		Escaper: `"`,
	}
}

// NewVesClient return a pointer of VesClinet
func NewVesClient(rOptions ...interface{}) (vc *VesClient, err error) {
	options := parseOptions(rOptions)
	vc = &VesClient{
		p:       newModelModule(),
		cb:      make(chan *bytes.Buffer, 1),
		quit:    make(chan bool, 1),
		logger:  options.logger,
		waitOpt: options.waitOpt,
		name:    options.clientName,
		nsbBase: options.nsbBase,

		module:    newDepModule(),
		nsbClient: nsbcli.NewNSBClient(options.nsbHost),
		dns:       xconfig.ChainDNS,
	}
	vc.module.Provide(config.ModulePath.Minimum.Global.Configuration, cfgX{
		dc: DatabaseConfig{DataFilePath: string(vc.name) + ".db"},
	})
	vc.module.Provide(config.ModulePath.Minimum.Global.Logger, vc.logger)
	if !vc.p.Install(vc.module.Module) {
		err = errInitModel
		return
	}
	if vc.db, err = NewAccountDB(vc.module); err != nil {
		return
	}
	vc.conn, _, err = new(websocket.Dialer).Dial((&url.URL{Scheme: "ws", Host: options.addr, Path: "/"}).String(), nil)
	return
}
