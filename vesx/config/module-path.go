package config

import (
	"github.com/Myriad-Dreamin/go-ves/lib/core"
	"path/filepath"
)

var (
	joiner           = filepath.Join
	base             = "ves"
	global           = joiner(base, "global")
	dbInstance       = joiner(base, "DBInstance")
	signer           = joiner(global, "Signer")
	centralVESClient = joiner(global, "CentralVESClient")
	nsbClient        = joiner(global, "NSBClient")
	respAccount      = joiner(global, "RespAccount")
	index            = joiner(dbInstance, "Index")
)

type GlobalPathS struct {
	CentralVESClient string
	Signer           string
	NSBClient        string
	RespAccount      string
}

type DBInstanceS struct {
	Index string
}

type ModulePathS struct {
	Minimum    mcore.ModulePathS
	Global     GlobalPathS
	DBInstance DBInstanceS
}

var ModulePath = ModulePathS{Minimum: mcore.DefaultNamespace,
	Global: GlobalPathS{
		CentralVESClient: centralVESClient,
		Signer:           signer,
		NSBClient:        nsbClient,
		RespAccount:      respAccount,
	},
	DBInstance: DBInstanceS{
		Index: index,
	},
}
