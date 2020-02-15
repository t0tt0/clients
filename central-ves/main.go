package main

import (
	"flag"
	centered_ves "github.com/Myriad-Dreamin/go-ves/central-ves/plugin/web-socket"
	"github.com/Myriad-Dreamin/go-ves/central-ves/server"
	_ "net/http/pprof"
)

var (
	port    = flag.String("port", ":23336", "serve on port")
	isDebug = flag.Bool("debug", false, "serve with debug mode")
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	srv := server.New("./config")
	if srv == nil {
		return
	}

	// srv.Inject(myPlugins...)
	if !srv.Inject(centered_ves.New()) {
		return
	}

	if *isDebug {
		srv.ServeWithPProf(*port)
	} else {
		srv.Serve(*port)
	}

}
