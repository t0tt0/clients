package index

import (
	"errors"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDBEngine struct {
	DB   *leveldb.DB
	wOpt *opt.WriteOptions
	rOpt *opt.ReadOptions
}

type LevelDBIndex struct {
	DB   *leveldb.DB
	wOpt *opt.WriteOptions
	rOpt *opt.ReadOptions
}

func (l LevelDBIndex) Get(k []byte) ([]byte, error) {
	return l.DB.Get(k, l.rOpt)
}

func (l LevelDBIndex) Set(k []byte, v []byte) error {
	return l.DB.Put(k, v, l.wOpt)
}

func (l LevelDBIndex) Delete(k []byte) error {
	return l.DB.Delete(k, l.wOpt)
}

func (l LevelDBIndex) Batch(ks [][]byte, vs [][]byte) error {
	return batch(l.DB, l.wOpt, ks, vs)
}

func (db *LevelDBEngine) ToIndex() types.Index {
	return LevelDBIndex{DB: db.DB, wOpt: db.wOpt, rOpt: db.rOpt}
}

func (db *LevelDBEngine) Get(k []byte) (ByteObject, error) {
	return ToByteObject(db.DB.Get(k, db.rOpt))
}

func (db *LevelDBEngine) Put(k []byte, v []byte) error {
	return db.DB.Put(k, v, db.wOpt)
}

func batch(db *leveldb.DB, options *opt.WriteOptions, ks, vs [][]byte) error {
	if len(ks) != len(vs) {
		return errors.New("inconsistent len of slice ks and vs")
	}
	var err error
	for i := range ks {
		if err = db.Put(ks[i], vs[i], options); err != nil {
			return err
		}
	}
	return nil
}

func (db *LevelDBEngine) Batch(ks [][]byte, vs [][]byte) error {
	return batch(db.DB, db.wOpt, ks, vs)
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
