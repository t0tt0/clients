package control

import "github.com/HyperService-Consortium/go-ves/grpc/uiprpc"

type SessionService interface {
	SessionServiceSignatureXXX() interface{}
	uiprpc.VESServer
}
