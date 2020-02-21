package bni

import (
	"github.com/HyperService-Consortium/go-ves/lib/upstream"
)

func (bn *BN) CheckAddress(addr []byte) error {
	if len(addr) == 20 {
		return nil
	} else {
		return upstream.ErrInvalidLength
	}
}
