package miris

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/kataras/iris"
)

type HandlerFunc = controller.HandlerFunc

func ToIrisHandler(h HandlerFunc) iris.Handler {
	return func(c iris.Context) {
		h(Context{Context: c})
		return
	}
}

func ToIrisHandlers(funcs []controller.HandlerFunc) (rFuncs []iris.Handler) {
	rFuncs = make([]iris.Handler, len(funcs))
	for i := range funcs {
		rFuncs[i] = ToIrisHandler(funcs[i])
	}
	return
}
