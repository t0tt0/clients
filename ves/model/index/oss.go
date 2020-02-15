package index

import (
	"errors"
	"github.com/Myriad-Dreamin/go-ves/types"
)

var engine Engine

var ErrNotExist = errors.New("object not exists")

type ByteObject interface {
	Data() []byte
	Free()
}

type byteObject []byte

func (b byteObject) Data() []byte { return b }
func (b byteObject) Free()        {}

func ToByteObject(obj []byte, err error) (ByteObject, error) {
	if err != nil {
		return nil, err
	}
	return byteObject(obj), nil
}

type Engine interface {
	Get([]byte) (ByteObject, error)
	Put([]byte, []byte) error
	Delete([]byte) error
	//Batch([][]byte, [][]byte) error
	Close() error
}

type ToIndexInterface interface {
	ToIndex() types.Index
}

func ToIndex(e Engine) types.Index {
	if x, ok := e.(ToIndexInterface); ok {
		return x.ToIndex()
	}
	return nil
}


func RegisterEngine(e Engine) error {
	engine = e
	return nil
}
