package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"google.golang.org/grpc"
)

type vesConn struct {
	conn *grpc.ClientConn
	uiprpc.VESClient
}

func (vc *VesClient) getVESConn(targetHost string, c *vesConn) (err error) {
	c.conn, err = grpc.Dial(targetHost, grpc.WithInsecure())
	if err != nil {
		return wrapper.Wrap(types.CodeNotConnected, err)
	}
	c.VESClient = uiprpc.NewVESClient(c.conn)
	return
}

func (vc *VesClient) ensureVESConn(targetHost string, c *vesConn) bool {
	err := vc.getVESConn(targetHost, c)
	if err != nil {
		vc.logger.Error("get ves conn error", "error", err)
		return false
	}
	return true
}

func (vc *VesClient) putVESConn(c *vesConn) error {
	return c.conn.Close()
}

func (vc *VesClient) mustPutVESConn(c *vesConn) bool {
	err := vc.putVESConn(c)
	if err != nil {
		vc.logger.Error("put ves conn error", "error", err)
		return false
	}
	return true
}
