package vesclient

import (
	"context"
	"fmt"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
)

func (vc *VesClient) SendOpIntentsByStrings(
	targetHost string, intents []string, deps []string) ([]byte, error) {
	return vc.SendOpIntents(
		targetHost,
		stringSliceToBytesSlice(intents),
		stringSliceToBytesSlice(deps))
}

func (vc *VesClient) SendOpIntentsR(
	targetHost string, content []byte) ([]byte, error) {
	var c vesConn
	vc.ensureVESConn(targetHost, &c)
	fmt.Println(targetHost)

	ctx, cancel := context.WithTimeout(
		context.Background(), vc.constant.SendOpIntentsTimeout)
	defer cancel()
	r, err := c.SessionStartR(
		ctx,
		&uiprpc.SessionStartRequestR{
			Content: content,
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
