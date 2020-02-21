package config

import (
	"github.com/Myriad-Dreamin/go-ves/lib/backend/core"
	"path/filepath"
)

var (
	joiner = filepath.Join
	base   = "ves"

	global     = joiner(base, "Global")
	dbInstance = joiner(base, "DBInstance")
	service    = joiner(base, "Service")

	signer           = joiner(global, "Signer")
	router           = joiner(global, "Router")
	centralVESClient = joiner(global, "CentralVESClient")
	nsbClient        = joiner(global, "NSBClient")
	respAccount      = joiner(global, "RespAccount")
	storage          = joiner(global, "Storage")
	storageHandler   = joiner(global, "StorageHandler")
	loggerWriter     = joiner(global, "LoggerWriter")
	closeHandler     = joiner(global, "CloseHandler")

	index       = joiner(dbInstance, "Index")
	modelModule = joiner(dbInstance, "ModelModule")

	vesServer           = joiner(service, "VesServer")
	chainDNS            = joiner(service, "ChainDNS")
	opIntentInitializer = joiner(service, "OpIntentInitializer")
)

type GlobalPathS struct {
	CentralVESClient string
	Signer           string
	NSBClient        string
	RespAccount      string
	Storage          string
	StorageHandler   string
	LoggerWriter     string
	Router           string
	CloseHandler     string
}

type DBInstanceS struct {
	Index       string
	ModelModule string
}

type ServiceS struct {
	VESServer           string
	ChainDNS            string
	OpIntentInitializer string
}

type ModulePathS struct {
	Minimum    mcore.ModulePathS
	Global     GlobalPathS
	DBInstance DBInstanceS
	Service    ServiceS
}

var ModulePath = ModulePathS{Minimum: mcore.DefaultNamespace,
	Global: GlobalPathS{
		LoggerWriter:     loggerWriter,
		CentralVESClient: centralVESClient,
		Signer:           signer,
		NSBClient:        nsbClient,
		RespAccount:      respAccount,
		Storage:          storage,
		StorageHandler:   storageHandler,
		Router:           router,
		CloseHandler:     closeHandler,
	},
	DBInstance: DBInstanceS{
		Index:       index,
		ModelModule: modelModule,
	},

	Service: ServiceS{
		VESServer:           vesServer,
		ChainDNS:            chainDNS,
		OpIntentInitializer: opIntentInitializer,
	},
}
