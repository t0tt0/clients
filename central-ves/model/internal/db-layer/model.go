package dblayer

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
)

type ChainInfo = database.ChainInfo
type Object = database.Object
type User = database.User

type ChainInfoFilter = database.ChainInfoFilter

func encodeAddress(src []byte) string {
	return database.EncodeAddress(src)
}

func decodeAddress(src string) ([]byte, error) {
	return database.DecodeAddress(src)
}
