package bni

import (
	"github.com/Myriad-Dreamin/go-ves/lib/bni/upstream"
)

func (bn *BN) CheckAddress(addr []byte) error {
	if len(addr) == 20 {
		return nil
	} else {
		return upstream.ErrInvalidLength
	}
}
