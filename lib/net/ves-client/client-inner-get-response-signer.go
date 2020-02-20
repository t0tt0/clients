package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/errorc"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
)

func (vc *VesClient) getRespSigner(acc uip.Account) (uip.Signer, error) {
	rawAcc, err := vc.db.InvertFind(acc)
	if errS := errorc.MaybeSelectError(rawAcc, err); errS.Code != 0 {
		return nil, wrapper.Wrap(errS.Code, errS)
	}

	return vc.AccountToSigner(rawAcc)
}
