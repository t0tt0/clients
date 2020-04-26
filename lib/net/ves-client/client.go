package vesclient

import (
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/go-uip/uip"
	xconfig "github.com/HyperService-Consortium/go-ves/config"
	core_cfg "github.com/HyperService-Consortium/go-ves/lib/backend/core-cfg"
	"github.com/HyperService-Consortium/go-ves/lib/net/ves-websocket"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/config"
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
	quit                   chan bool

	db        AccountDBInterface
	sessionDB SessionDBInterface
	conn      ves_websocket.VESWSSocket
	nsbSigner uip.Signer
	dns       types.ChainDNSInterface
	nsbClient types.NSBClient

	waitOpt              uip.RouteOptionTimeout
	name                 []byte
	ignoreUnknownMessage bool
	nsbBase              string

	// client scope default nsbHost, vesHost
	nsbHost  string
	vesHost  string
	constant *ClientConstant
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

func (vc *VesClient) closeHandler(code int, text string) error {
	if code != websocket.CloseNoStatusReceived {
		vc.logger.Info("closed", "code", code, "text", text)
	}
	return nil
}

// NewVesClient return a pointer of VesClient
func NewVesClient(rOptions ...interface{}) (vc *VesClient, err error) {
	options := parseOptions(rOptions)
	vc = &VesClient{
		quit:     make(chan bool, 1),
		logger:   options.logger,
		name:     options.clientName,
		nsbBase:  options.nsbBase,
		waitOpt:  options.waitOpt,
		constant: options.constant,

		p:         newModelModule(),
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
	if vc.sessionDB, err = NewSessionDB(vc.module); err != nil {
		return
	}
	var conn ves_websocket.SocketConn
	conn, _, err = new(websocket.Dialer).Dial((&url.URL{Scheme: "ws", Host: options.addr, Path: "/"}).String(), nil)
	if err != nil {
		return
	}
	vc.conn, err = ves_websocket.NewVESSocket(conn, vc.ProcessMessage, vc.logger)
	vc.conn.SetCloseHandler(vc.closeHandler)
	return
}
