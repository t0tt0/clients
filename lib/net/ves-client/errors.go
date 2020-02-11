package vesclient

import "errors"

var (
	errIlegalPrivateKey = errors.New("illegal private key")
	errTen1NotFound     = errors.New("key ten1 not found, which is used for communicating with nsb")
	errNotFound         = errors.New("not found")
	errInitModel = errors.New("init model error")
)
