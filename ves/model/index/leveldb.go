package index

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDBEngine struct {
	DB   *leveldb.DB
	wOpt *opt.WriteOptions
	rOpt *opt.ReadOptions
}

func (db *LevelDBEngine) Get(k []byte) (ByteObject, error) {
	return ToByteObject(db.DB.Get(k, db.rOpt))
}

func (db *LevelDBEngine) Put(k []byte, v []byte) error {
	return db.DB.Put(k, v, db.wOpt)
}

func (db *LevelDBEngine) Delete(k []byte) error {
	return db.DB.Delete(k, db.wOpt)
}

func (db *LevelDBEngine) Close() error {
	return db.DB.Close()
}

func NewLevelDB(path string, opts *opt.Options) (Engine, error) {
	e := new(LevelDBEngine)
	var err error
	e.DB, err = leveldb.OpenFile(path, opts)
	if err != nil {
		return nil, err
	}
	e.wOpt = &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         true,
	}
	e.rOpt = &opt.ReadOptions{
		DontFillCache: false,
		Strict:        0,
	}
	return e, nil
}
