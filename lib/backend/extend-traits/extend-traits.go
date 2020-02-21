package extend_traits

import (
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type ExtendModelOperationInterface = traits.Interface

type ExtendModel struct {
	i           ExtendModelOperationInterface
	replacement interface{}
}

func NewExtendModel(t ExtendModelOperationInterface) ExtendModel {
	return ExtendModel{
		i:           t,
		replacement: t.ObjectFactory(),
	}
}
