package control

import "github.com/Myriad-Dreamin/minimum-lib/module"

type Dependencies = module.Module
type Engine interface {
	Build(m Dependencies) error
	RunnableEngine
}

type RunnableEngine interface {
	Run(port string) error
}
