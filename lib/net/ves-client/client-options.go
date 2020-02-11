package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"time"
)

type CVesHostOption string
type NsbHostOption string
type VesName []byte

type ServerOptions struct {
	logger  logger.Logger
	waitOpt uiptypes.RouteOptionTimeout
	addr    string
	nsbHost string
	vesName []byte
}

var globalLogger = logger.NewStdLogger()

func defaultServerOptions() ServerOptions {
	return ServerOptions{
		logger:  globalLogger,
		waitOpt: uiptypes.RouteOptionTimeout(time.Second * 60),
		addr:    "127.0.0.1:23452",
		nsbHost: "127.0.0.1:27667",
	}
}

func parseOptions(rOptions []interface{}) ServerOptions {
	var options = defaultServerOptions()
	for i := range rOptions {
		switch option := rOptions[i].(type) {
		case logger.Logger:
			options.logger = option
		case uiptypes.RouteOptionTimeout:
			options.waitOpt = option
		case CVesHostOption:
			options.addr = string(option)
		case NsbHostOption:
			options.nsbHost = string(option)
		case VesName:
			options.vesName = option
		}
	}
	return options
}


