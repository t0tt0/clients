package model

import (
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/provider"
)

type Provider = abstraction.Provider

func NewProvider(namespace string) Provider {
	return provider.NewProvider(namespace)
}
