package vesclient

import (
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
)

func (vc *VesClient) AccountToSigner(account *Account) (signer uip.Signer, err error) {
	addition, err := decodeAddition(account.Addition)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeDecodeAdditionError, err)
	}
	switch ChainType.Type(account.ChainType) {
	case ChainType.Ethereum:
		address, err := decodeAddition(account.Address)
		if err != nil {
			return nil, wrapper.Wrap(types.CodeDecodeAddressError, err)
		}
		signer, err = NewEthAccount(address, addition, account.ChainID)
	case ChainType.TendermintNSB:
		signer, err = signaturer.NewTendermintNSBSigner(addition)
	default:
		return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
	}
	if err != nil {
		return nil, wrapper.Wrap(types.CodeConvertSignerError, err)
	}
	return signer, nil
}
