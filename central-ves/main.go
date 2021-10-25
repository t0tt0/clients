package main

import (
	"flag"
	"github.com/HyperService-Consortium/go-ves/central-ves/server"
	"github.com/HyperService-Consortium/go-ves/central-ves/web-socket"
	_ "net/http/pprof"
)

var (
	port    = flag.String("port", ":23336", "serve on port")
	config  = flag.String("config", "./config", "config file path")
	isDebug = flag.Bool("debug", false, "serve with debug mode")
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
//config: config file path
	srv := server.New(*config)
	if srv == nil {
		return
	}

	// srv.Inject(myPlugins...)
//New(): new plugin for the websocket server
	if !srv.Inject(central_ves.New()) {
		return
	}

	if *isDebug {
		srv.ServeWithPProf(*port)
	} else {
		srv.Serve(*port)
	}

}
