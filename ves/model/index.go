package model

import (
	"github.com/HyperService-Consortium/go-ves/ves/model/index"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

//func NewRocksDB(opts *gorocksdb.Options, name string) (oss.Engine, error) {
//	return oss.NewRocksDB(opts, name)
//}

func NewLevelDB(path string, opts *opt.Options) (index.Engine, error) {
	return index.NewLevelDB(path, opts)
}

func RegisterIndex(e index.Engine) error {
	return index.RegisterEngine(e)
}
