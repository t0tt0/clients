package vesclient

import (
	"flag"
	"log"
)

func Init() {
	StartDaemon()
	if !flag.Parsed() {
		flag.Parse()
	}
}

// Main is the origin main of ves client
func Main(name string, addr string, port string) {
//VanillaMakeClient returns a web socket client
	vcClient, err := VanillaMakeClient(name, addr)
	if err != nil {
		log.Fatal("make client error", err)
	}

	if err = vcClient.Boot(); err != nil {
		log.Fatal("boot error", err)
	}

	if err = vcClient.ListenHTTP(port); err != nil {
		vcClient.logger.Fatal("listen error", "error", err)
	}
}
