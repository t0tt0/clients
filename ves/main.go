package main

import (
	"flag"
	"github.com/Myriad-Dreamin/go-ves/ves/server"
	"log"
	_ "net/http/pprof"
)

var (
	httpPort  = flag.String("port", ":23335", "serve http on port")
	gRPCPport = flag.String("grpc", ":23351", "serve grpc on port")
	config    = flag.String("config", "./config", "config file path")
	isDebug   = flag.Bool("debug", false, "serve with debug mode")
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	srv, err := server.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	// srv.Inject(myPlugins...)
	//httpPort, gRPCPort
	if *isDebug {
		srv.ServeWithPProf(*httpPort, *gRPCPport)
	} else {
		srv.Serve(*httpPort, *gRPCPport)
	}

}
