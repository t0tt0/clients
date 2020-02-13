package vesclient

import (
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"

	ethbni "github.com/Myriad-Dreamin/go-ves/lib/bni/eth"
	nsbbni "github.com/Myriad-Dreamin/go-ves/lib/bni/ten"
)

func (vc *VesClient) ensureRouter(chainID uint64, router *uiptypes.Router) bool {
	if *router != nil {
		return true
	}
	var err error
	if *router, err = vc.getRouter(chainID); err != nil {
		vc.logger.Error("get router error", "error", err)
	}
	return err == nil
}

func (vc *VesClient) getRouter(chainID uint64) (uiptypes.Router, error) {
	if ci, err := vc.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(vc.dns), nil
		case ChainType.TendermintNSB:
			return nsbbni.NewBN(vc.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}

func (vc *VesClient) getBlockStorage(chainID uint64) (uiptypes.Storage, error) {
	if ci, err := vc.dns.GetChainInfo(chainID); err != nil {
		return nil, err
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(vc.dns), nil
		case ChainType.TendermintNSB:
			return nsbbni.NewBN(vc.dns), nil
		default:
			return nil, wrapCode(CodeUnknownChainID)
		}
	}
}
