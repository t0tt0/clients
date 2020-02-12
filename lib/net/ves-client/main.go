package vesclient

import (
	"flag"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	name = flag.String("name", "test", "name of client, which must be provided")
	addr = flag.String("addr", "localhost:23452", "http service address")
)

func Init() {
	StartDaemon()
	if !flag.Parsed() {
		flag.Parse()
	}
}

// Main is the origin main of ves client
func Main() {
	vesLogger, err := logger.NewZapLogger(
		logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
	if err != nil {
		log.Fatal("init vesLogger error", "error", err)
	}
	vcClient, err := NewVesClient(vesLogger, ClientName(*name), CVesHostOption(*addr))
	if err != nil {
		vesLogger.Fatal("get   ves client error", "error", err)
	}
	if err = vcClient.ListenHTTP(); err != nil {
		vesLogger.Fatal("listen error", "error", err)
	}

	//fmt.Println("input your name:")
	//vcClient.name, _, err = bufio.NewReader(os.Stdin).ReadLine()
	//if err != nil {
	//	vcClient.vesLogger.Fatal("set name", "error", err)
	//}
	//
	//if err = vcClient.Boot(); err != nil {
	//	os.Exit(1)
	//}
	//
	//phandler.register(func() { vcClient.quit <- true })
	//go vcClient.write()
	//select {
	//case <-vcClient.quit:
	//	return
	//}
}
