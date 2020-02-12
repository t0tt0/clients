package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
)

func (vc *VesClient) ensureGetNSBSigner(signer *uiptypes.Signer) bool {
	if *signer != nil {
		return true
	}
	var err error
	if *signer, err = vc.getNSBSigner(); err != nil {
		vc.logger.Error("get nsb signer error", "error", err)
		return false
	} else {
		return true
	}

}

func (vc *VesClient) getNSBSigner() (uiptypes.Signer, error) {
	if vc.signer != nil {
		return vc.signer, nil
	}

	key, err := vc.db.QueryAlias(vc.nsbBase)
	if err != nil {
		return nil, wrap(CodeSelectError, err)
	} else if key == nil {
		return nil, wrapCode(CodeNotFound)
	}

	b, err := decodeAddress(key.Address)
	if err != nil {
		return nil, wrap(CodeDecodeAddressError, err)
	}
	vc.signer, err = signaturer.NewTendermintNSBSigner(b)
	if err != nil {
		return nil, wrap(CodeInitializeNSBSignerError, err)
	}

	return vc.signer, nil
}
