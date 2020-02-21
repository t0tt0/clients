package splayer

import (
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type ObjectDB struct {
	abstraction.ObjectDB
}

func NewObjectDB(base abstraction.ObjectDB, m module.Module) (*ObjectDB, error) {
	db := new(ObjectDB)
	db.ObjectDB = base
	return db, nil
}
