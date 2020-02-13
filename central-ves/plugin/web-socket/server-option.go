package centered_ves

import "github.com/Myriad-Dreamin/minimum-lib/logger"

type NSBHostOption string
type ServerOptions struct {
	logger  logger.Logger
	nsbHost string
}

func defaultServerOptions() ServerOptions {
	return ServerOptions{
		logger:  logger.NewStdLogger(),
		nsbHost: "127.0.0.1:26657",
	}
}

func parseOptions(rOptions []interface{}) ServerOptions {
	var options = defaultServerOptions()
	for i := range rOptions {
		switch option := rOptions[i].(type) {
		case logger.Logger:
			options.logger = option
		case NSBHostOption:
			options.nsbHost = string(option)
		}
	}
	return options
}

