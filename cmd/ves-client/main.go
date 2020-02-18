package main

import (
	"flag"
	vesclient "github.com/Myriad-Dreamin/go-ves/lib/net/ves-client"
)

var (
	name = flag.String("name", "test", "name of client, which must be provided")
	addr = flag.String("addr", "localhost:23452", "http service address")
	port = flag.String("port", ":26670", "listening http Port")
)

func init() {
	vesclient.Init()
}

func main() {
	vesclient.Main(*name, *addr, *port)
}
