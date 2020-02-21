package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
)

func (vc *VesClient) ensureGetNSBSigner(signer *uip.Signer) bool {
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

func (vc *VesClient) getNSBSigner() (uip.Signer, error) {
	if vc.nsbSigner != nil {
		return vc.nsbSigner, nil
	}

	key, err := vc.db.QueryAlias(vc.nsbBase)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSelectError, err)
	} else if key == nil {
		return nil, wrapper.WrapCode(types.CodeNotFound)
	}

	b, err := decodeAddition(key.Addition)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeDecodeAdditionError, err)
	}
	vc.nsbSigner, err = signaturer.NewTendermintNSBSigner(b)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeConvertSignerError, err)
	}

	return vc.nsbSigner, nil
}
