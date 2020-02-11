package vesclient

import "github.com/Myriad-Dreamin/go-ves/lib/wrapper"

type Code int

const (
	CodeOk Code = iota
	CodeSelectError
	CodeNotFound

	CodeDecodeAddressError
	CodeInitializeNSBSignerError

	CodeUnknownChainID
)

func wrap(code Code, err error) error {
	return wrapper.WrapN(3, int(code), err)
}

func wrapCode(code Code) error {
	return wrapper.WrapCodeN(3, int(code))
}
