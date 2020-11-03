package vesclient

import (
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"go.uber.org/zap/zapcore"
	"log"
)

// VanillaMakeClient for humans
func VanillaMakeClient(name, addr string, options ...interface{}) (*VesClient, error) {
	vesLogger, err := logger.NewZapLogger(
		logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
	if err != nil {
		log.Fatal("init vesLogger error", "error", err)
	}
	vesLogger.With("client-name", name)
	//NewVesClient: define processmessage which sends and receives requests from other peers' cservers and servers
	vcClient, err := NewVesClient(vesLogger, ClientName(name), CVesHostOption(addr))
	if err != nil {
		vesLogger.Fatal("get ves client error", "error", err)
	}

	return vcClient, nil
}
