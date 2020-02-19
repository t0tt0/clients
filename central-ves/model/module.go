package model

import (
	mcore "github.com/Myriad-Dreamin/go-ves/lib/backend/core"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Module struct {
	dbLayer DBLayerModule
	spLayer SPLayerModule
}

func NewModule() Module {
	return Module{
		dbLayer: NewDBLayerModule(),
		spLayer: NewSPLayerModule(),
	}
}

func (p *Module) InstallFromContext(dep module.Module) bool {
	return p.dbLayer.FromContext(dep) && p.spLayer.FromContext(dep)
}

func (p *Module) InstallFromConfiguration(dep module.Module) bool {
	return p.dbLayer.InstallFromConfiguration(dep) && p.spLayer.InstallFromConfiguration(dep)
}

func (p *Module) Install(dep module.Module) bool {
	return p.dbLayer.Install(dep) && p.spLayer.Install(dep)
}

func (p *Module) InstallMock(dep module.Module, callback mcore.MockCallback) bool {
	return p.dbLayer.InstallMock(dep, callback)
}

func (p *Module) Close(dep module.Module) bool {
	return p.dbLayer.Close(dep) && p.spLayer.Close(dep)
}
