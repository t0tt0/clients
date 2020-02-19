package bni

import "github.com/Myriad-Dreamin/go-ves/lib/upstream"

func (bn *BN) CheckAddress(addr []byte) error {
	if len(addr) == 32 {
		return nil
	} else {
		return upstream.ErrInvalidLength
	}
}
