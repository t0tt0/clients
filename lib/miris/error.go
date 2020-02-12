package miris

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type IrisError struct {
	Err  error
	Type controller.ErrorType
	Meta interface{}
}

func (g IrisError) GetError() error {
	return g.Err
}

func (g IrisError) GetType() controller.ErrorType {
	return uint64(g.Type)
}

func (g IrisError) GetMeta() interface{} {
	return g.Meta
}
