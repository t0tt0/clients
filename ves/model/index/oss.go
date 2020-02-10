package index

import (
	"errors"
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
	Close() error
}

func RegisterEngine(e Engine) error {
	engine = e
	return nil
}
