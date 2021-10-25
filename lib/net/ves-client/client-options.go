package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"time"
)

type CVesHostOption string
type NsbHostOption string
type ClientName []byte

type ClientConstant struct {
	SendOpIntentsTimeout time.Duration
}

type ServerOptions struct {
	logger     logger.Logger
	waitOpt    uip.RouteOptionTimeout
	addr       string
	nsbHost    string
	nsbBase    string
	clientName []byte
	constant   *ClientConstant
}

var globalLogger = logger.NewStdLogger()

func NewConstantOption() *ClientConstant {
	return &ClientConstant{
		SendOpIntentsTimeout: time.Minute,
	}
}

func defaultServerOptions() ServerOptions {
	return ServerOptions{
		logger:     globalLogger,
		waitOpt:    uip.RouteOptionTimeout(time.Second * 60),
		clientName: []byte("test"),
		addr:       "127.0.0.1:23452",
		nsbBase:    "ten1",
		nsbHost:    "127.0.0.1:27667",
		constant:   NewConstantOption(),
	}
}

func parseOptions(rOptions []interface{}) ServerOptions {
	var options = defaultServerOptions()
	for i := range rOptions {
		switch option := rOptions[i].(type) {
		case logger.Logger:
			options.logger = option
		case uip.RouteOptionTimeout:
			options.waitOpt = option
		case CVesHostOption:
			options.addr = string(option)
		case NsbHostOption:
			options.nsbHost = string(option)
		case ClientName:
			options.clientName = option
		case *ClientConstant:
			options.constant = option
		}
	}
	return options
}
