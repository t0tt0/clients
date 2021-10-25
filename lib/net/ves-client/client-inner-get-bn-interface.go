package vesclient

import (
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"

	ethbni "github.com/HyperService-Consortium/go-ves/lib/bni/eth"
	nsbbni "github.com/HyperService-Consortium/go-ves/lib/bni/ten"
)
//from websocket information, need to know which chain, not initializing ops
func (vc *VesClient) ensureRouter(chainID uint64, router *uip.Router) bool {
	if *router != nil {
		return true
	}
	var err error
	if *router, err = vc.getRouter(chainID); err != nil {
		vc.logger.Error("get router error", "error", err)
	}
	return err == nil
}

func (vc *VesClient) getRouter(chainID uint64) (uip.Router, error) {
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

func (vc *VesClient) ensureBlockStorage(chainID uint64, storage *uip.Storage) bool {
	if *storage != nil {
		return true
	}
	var err error
	if *storage, err = vc.getBlockStorage(chainID); err != nil {
		vc.logger.Error("get storage error", "error", err)
	}
	return err == nil
}

func (vc *VesClient) getBlockStorage(chainID uint64) (uip.Storage, error) {
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

func (vc *VesClient) ensureTranslator(chainID uint64, storage *uip.Translator) bool {
	if *storage != nil {
		return true
	}
	var err error
	if *storage, err = vc.getTranslator(chainID); err != nil {
		vc.logger.Error("get translator error", "error", err)
	}
	return err == nil
}

func (vc *VesClient) getTranslator(chainID uint64) (uip.Translator, error) {
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
