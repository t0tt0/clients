package sessionservice

import (
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	ethbni "github.com/HyperService-Consortium/go-ves/lib/bni/eth"
	tenbni "github.com/HyperService-Consortium/go-ves/lib/bni/ten"
	"github.com/HyperService-Consortium/go-ves/types"
)

func (svc *Service) getBlockChainInterface(chainID uint64) (uip.BlockChainInterface, error) {
	if ci, err := svc.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(svc.dns), nil
		case ChainType.TendermintNSB:
			return tenbni.NewBN(svc.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}
