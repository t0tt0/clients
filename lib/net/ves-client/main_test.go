package vesclient

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"os"
	"testing"
)

var describer = wrapper.Describer{
	Pack: "github.com/Myriad-Dreamin/go-ves/lib/net/ves-client",
	Rel:  sugar.HandlerError(os.Getwd()).(string)}

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

func testInit() {
	wrapper.SetCodeDescriptor(func(code int) string {
		return types.CodeType(code).String()
	})
}

func TestMain(m *testing.M) {
	testInit()
	m.Run()
}
