package vesclient

import (
	"bytes"
	"errors"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"os"
	"sync"
	"testing"
)

var describer = wrapper.Describer{
	Pack: "github.com/Myriad-Dreamin/go-ves/lib/net/ves-client",
	Rel:  sugar.HandlerError(os.Getwd()).(string)}

type fields struct {
	p                      modelModule
	rwMutex                sync.RWMutex
	logger                 logger.Logger
	module                 DepModule
	closeSessionRWMutex    sync.RWMutex
	closeSessionSubscriber []SessionCloseSubscriber
	name                   []byte
	db                     AccountDBInterface
	conn                   SocketConn
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

type fieldOption struct {
	dns types.ChainDNSInterface
}

type ChainDNSMockData struct {
	K uiptypes.ChainIDUnderlyingType
	V ChainInfo
}

type ChainInfo struct {
	ChainType uiptypes.ChainType
	ChainHost string
}

func (c ChainInfo) GetChainType() uiptypes.ChainType {
	return c.ChainType
}

func (c ChainInfo) GetChainHost() string {
	return c.ChainHost
}

type mockChainDNSImpl map[uiptypes.ChainIDUnderlyingType]types.ChainInfo

func (m mockChainDNSImpl) GetChainInfo(
	chainID uiptypes.ChainIDUnderlyingType) (types.ChainInfo, error) {
	ci, ok := m[chainID]
	if !ok {
		return nil, errors.New("not found")
	}
	return ci, nil
}

func mockChainDNS(data ...ChainDNSMockData) types.ChainDNSInterface {
	dns := make(mockChainDNSImpl)
	for _, ci := range data {
		dns[ci.K] = ci.V
	}
	return dns
}

func parseFieldOptions(rawOpts []interface{}) (options fieldOption) {
	options = fieldOption{
		dns: nil,
	}
	for i := range rawOpts {
		switch o := rawOpts[i].(type) {
		case fieldOptionDNS:
			options.dns = o
		}
	}
	return
}

func createFields(rawOpts ...interface{}) fields {
	options := parseFieldOptions(rawOpts)
	return fields{
		p:                      newModelModule(),
		rwMutex:                sync.RWMutex{},
		logger:                 nil,
		module:                 DepModule{},
		closeSessionRWMutex:    sync.RWMutex{},
		closeSessionSubscriber: nil,
		name:                   nil,
		db:                     nil,
		conn:                   nil,
		nsbSigner:              nil,
		dns:                    options.dns,
		nsbClient:              nil,
		waitOpt:                0,
		cb:                     nil,
		quit:                   nil,
		nsbip:                  "",
		grpcip:                 "",
		nsbBase:                "",
	}
}

func testInit() {
	wrapper.SetCodeDescriptor(func(code int) string {
		return types.CodeType(code).String()
	})
}

func TestMain(m *testing.M) {
	testInit()
	m.Run()
}
