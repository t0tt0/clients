package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
)

type Code int

const (
	CodeOk Code = iota
	CodeSelectError
	CodeNotFound

	CodeDecodeAddressError
	CodeInitializeNSBSignerError
	CodeReadMessageError
	CodeReadMessageIDError

	CodeUnknownChainID
)

func wrap(code Code, err error) error {
	return wrapper.WrapN(3, int(code), err)
}

func wrapCode(code Code) error {
	return wrapper.WrapCodeN(3, int(code))
}

func errorSerializer(code types.Code, err string) serial.ErrorSerializer {
	return serial.ErrorSerializer{
		Code: int(code),
		Err:  err,
	}
}

type response struct {
	Code int `json:"code"`
}

var responseOK = response{Code: int(CodeOk)}
