package vesclient

import (
	"context"
	"fmt"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
)

func (vc *VesClient) SendOpIntentsByStrings(
	targetHost string, intents []string, deps []string) ([]byte, error) {
	return vc.SendOpIntents(
		targetHost,
		stringSliceToBytesSlice(intents),
		stringSliceToBytesSlice(deps))
}

func (vc *VesClient) SendOpIntents(
	targetHost string, intents [][]byte, deps [][]byte) ([]byte, error) {
	var c vesConn
	vc.ensureVESConn(targetHost, &c)
	fmt.Println(targetHost)

	ctx, cancel := context.WithTimeout(
		context.Background(), vc.constant.SendOpIntentsTimeout)
	defer cancel()
	r, err := c.SessionStart(
		ctx,
		&uiprpc.SessionStartRequest{
			Opintents: &uiprpc_base.OpIntents{
				Dependencies: deps,
				Contents:     intents,
			},
		})
	vc.mustPutVESConn(&c)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeExecuteError, err)
	}
	if !r.GetOk() {
		return nil, wrapper.WrapCode(types.CodeExecuteError)
	}

	return r.GetSessionId(), nil
}
