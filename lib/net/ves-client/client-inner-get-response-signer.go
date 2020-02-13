package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/lib/errorc"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
)

func (vc *VesClient) getRespSigner(acc *uiprpc_base.Account) (uiptypes.Signer, error) {
	rawAcc, err := vc.db.InvertFind(acc)
	if errS := errorc.MaybeSelectError(rawAcc, err); errS.Code != 0 {
		return nil, wrapper.Wrap(errS.Code, errS)
	}

	return vc.AccountToSigner(rawAcc)
}
