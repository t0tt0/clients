package getter

import (
	"fmt"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	ethbni "github.com/HyperService-Consortium/go-ves/lib/bni/eth"
	tenbni "github.com/HyperService-Consortium/go-ves/lib/bni/ten"
	"github.com/HyperService-Consortium/go-ves/types"
)

type BlockChainGetter struct {
	dns types.ChainDNSInterface
}

func NewBlockChainGetter(dns types.ChainDNSInterface) *BlockChainGetter {
	return &BlockChainGetter{dns: dns}
}

func (b BlockChainGetter) GetChecker(chainID uip.ChainID) (checker uip.Checker, err error) {
	if ci, err := b.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(b.dns), nil
		case ChainType.TendermintNSB:
			return tenbni.NewBN(b.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}

func (b BlockChainGetter) GetTranslator(chainID uip.ChainID) (intent uip.Translator, err error) {
	if ci, err := b.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(b.dns), nil
		case ChainType.TendermintNSB:
			return tenbni.NewBN(b.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}

func (b BlockChainGetter) GetRouter(chainID uip.ChainID) (router uip.Router, err error) {
	if ci, err := b.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(b.dns), nil
		case ChainType.TendermintNSB:
			return tenbni.NewBN(b.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}

func (b BlockChainGetter) GetBlockStorage(chainID uip.ChainID) (storage uip.Storage, err error) {
	if ci, err := b.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(b.dns), nil
		case ChainType.TendermintNSB:
			return tenbni.NewBN(b.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}

//todo
func (b BlockChainGetter) GetBlockChainInterface(chainID uip.ChainID) uip.BlockChainInterface {
	if ci, err := b.dns.GetChainInfo(chainID); err != nil {
		fmt.Println(wrapper.Wrap(types.CodeChainIDNotFound, err))
		return nil
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(b.dns)
		case ChainType.TendermintNSB:
			return tenbni.NewBN(b.dns)
		default:
			fmt.Println(wrapper.WrapCode(types.CodeChainTypeNotFound))
			return nil
		}
	}
}
