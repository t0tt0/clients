package control

import "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"

type SessionService interface {
	SessionServiceSignatureXXX() interface{}
	uiprpc.VESServer
}
