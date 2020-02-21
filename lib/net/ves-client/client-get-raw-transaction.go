package vesclient

import (
	"context"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
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
