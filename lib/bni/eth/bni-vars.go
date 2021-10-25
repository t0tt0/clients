package bni

import "errors"

var (
	ErrNotSigned       = errors.New("not signed raw transaction")
	ErrDeployFailed    = errors.New("deploy failed")
	ErrTimeout         = errors.New("timeout")
	ErrNotMatchAddress = errors.New("not match address")
	ErrHasNoChainInfo  = errors.New("has no chain info")
)
