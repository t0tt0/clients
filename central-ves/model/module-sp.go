package model

import (
	mcore "github.com/Myriad-Dreamin/go-ves/lib/core"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type SPLayerModule struct {
	mcore.RedisPoolModule
	mcore.LoggerModule
	Opened bool
}

func NewSPLayerModule() SPLayerModule {
	return SPLayerModule{
		Opened: false,
	}
}

func (p *SPLayerModule) FromContext(dep module.Module) bool {
	p.Opened = p.LoggerModule.Install(dep) &&
		p.RedisPoolModule.FromContext(dep)
	return p.Opened
}

func (p *SPLayerModule) Install(dep module.Module) bool {
	p.Opened = p.LoggerModule.Install(dep) &&
		p.RedisPoolModule.InstallFromConfiguration(dep)
	return p.Opened
}

func (p *SPLayerModule) Close(dep module.Module) bool {
	if p.Opened {
		if err := p.RedisPool.Close(); err != nil {
			p.Logger.Error("close Redis error", "error", err)
			return false
		}
	}
	return true
}

