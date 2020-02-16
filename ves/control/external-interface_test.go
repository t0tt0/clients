package control

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/types"
)

var _ NSBClient = types.NSBClient(nil)
var _ types.NSBClient = NSBClient(nil)

var _ CentralVESClient = uiprpc.CenteredVESClient(nil)
var _ uiprpc.CenteredVESClient = CentralVESClient(nil)

var _ ChainDNS = types.ChainDNSInterface(nil)
var _ types.ChainDNSInterface = ChainDNS(nil)
