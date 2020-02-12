package vesclient

import "errors"

var (
	errIlegalPrivateKey = errors.New("illegal private key")
	errNotFound         = errors.New("not found")
	errInitModel        = errors.New("init model error")
)
