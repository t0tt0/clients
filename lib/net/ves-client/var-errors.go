package vesclient

import "errors"

var (
	errInitModel  = errors.New("init model error")
	errNilVesConn = errors.New("nil ves conn")
)
