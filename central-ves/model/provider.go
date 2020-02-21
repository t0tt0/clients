package model

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/provider"
)

type Provider = abstraction.Provider


func NewProvider(namespace string) Provider {
	return provider.NewProvider(namespace)
}
