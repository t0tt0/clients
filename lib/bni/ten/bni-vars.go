package bni

import "errors"

var (
	ErrorDecodeSrcAddress = errors.New("the src address should be in length of 32")
	ErrorDecodeDstAddress = errors.New("the dst address should be in length of 32 or 0")
	ErrNotSigned          = errors.New("not signed")
)
