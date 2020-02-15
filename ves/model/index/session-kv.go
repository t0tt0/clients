package index

import (
	"bytes"
	"github.com/Myriad-Dreamin/go-ves/types"
)

type SessionKV struct {
	i types.Index
}

func NewSessionKV(i types.Index) *SessionKV {
	return &SessionKV{i}
}

func decorate(p, b []byte) []byte {
	var x = bytes.NewBuffer(make([]byte, len(p)+2+len(b)))
	x.Reset()
	x.Write(p)
	x.WriteByte('k')
	x.WriteByte('k')
	x.Write(b)
	return x.Bytes()
}

type SessionFKV struct {
	i          *SessionKV
	iscAddress []byte
}

func (g *SessionFKV) Get(k []byte) ([]byte, error) {
	return g.i.GetKV(g.iscAddress, k)
}

func (g *SessionFKV) Set(k, v []byte) error {
	return g.i.SetKV(g.iscAddress, k, v)
}

func (g *SessionKV) GetKV(iscAddress, b []byte) ([]byte, error) {
	return g.i.Get(decorate(iscAddress, b))
}

func (g *SessionKV) SetKV(iscAddress, k, v []byte) error {
	return g.i.Set(decorate(iscAddress, k), v)
}
