package config

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/core"
	"path/filepath"
)

var (
	joiner              = filepath.Join
	base = "ves"
	global = joiner(base, "global")
	userDB = joiner(global, "userDB")
)

type GlobalPathS struct {
	UserDB string
}

type ModulePathS struct {
	Minimum mcore.ModulePathS
	Global GlobalPathS
}

var ModulePath = ModulePathS{Minimum:mcore.DefaultNamespace,
	Global: GlobalPathS{UserDB:userDB},
}
