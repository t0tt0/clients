package model

import (
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/fset"
)

type SessionFSet = fset.SessionFSet

func NewSessionFSet(p Provider, index types.Index) *SessionFSet {
	return fset.NewSessionFSet(p, index)
}
