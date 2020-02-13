package vs

import (
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
)

type VServer struct {
	Logger logger.Logger

	DB        model.VESDB
	Resp      *uiprpc_base.Account
	Signer    *signaturer.TendermintNSBSigner
	CVes      uiprpc.CenteredVESClient
	NsbClient *nsbcli.NSBClient
	Host      string
	NsbHost   string
}
