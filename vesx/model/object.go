package model

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	splayer "github.com/Myriad-Dreamin/go-ves/vesx/model/sp-layer"
)

type Object = splayer.Object
type ObjectDB = splayer.ObjectDB

func NewObjectDB(m module.Module) (*ObjectDB, error) {
	return splayer.NewObjectDB(m)
}

func GetObjectDB(m module.Module) (*ObjectDB, error) {
	return splayer.GetObjectDB(m)
}
