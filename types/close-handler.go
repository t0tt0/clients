package types

import "io"

type CloseHandler interface {
	io.Closer
	Handle(closer io.Closer)
}
