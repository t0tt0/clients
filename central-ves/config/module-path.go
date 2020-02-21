package config

import (
	"github.com/Myriad-Dreamin/go-ves/lib/backend/core"
	"path/filepath"
)

var (
	joiner       = filepath.Join
	base         = "ves"
	global       = joiner(base, "global")
	userDB       = joiner(global, "userDB")
	loggerWriter = joiner(global, "LoggerWriter")
	router       = joiner(global, "Router")
)

type GlobalPathS struct {
	UserDB       string
	LoggerWriter string
	Router       string
}

type ModulePathS struct {
	Minimum mcore.ModulePathS
	Global  GlobalPathS
}

var ModulePath = ModulePathS{Minimum: mcore.DefaultNamespace,
	Global: GlobalPathS{
		UserDB:       userDB,
		LoggerWriter: loggerWriter},
}
