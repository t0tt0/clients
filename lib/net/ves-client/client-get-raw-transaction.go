package vesclient

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
)

func (vc *VesClient) GetRawTransaction(sessionID []byte, targetHost string) (
	*uiprpc.SessionRequireRawTransactReply, error,
) {
	var c vesConn
	vc.ensureVESConn(targetHost, &c)

	ctx, cancel := context.WithTimeout(
		context.Background(), vc.constant.SendOpIntentsTimeout)
	defer cancel()
	r, err := c.SessionRequireRawTransact(
		ctx,
		&uiprpc.SessionRequireRawTransactRequest{
			SessionId: sessionID,
		},
	)
	vc.mustPutVESConn(&c)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeExecuteError, err)
	}
	return r, nil
}
