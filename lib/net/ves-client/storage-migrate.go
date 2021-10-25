package vesclient

import "github.com/HyperService-Consortium/go-ves/lib/basic/fcg"

func (m *modelModule) Migrates() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//migrations
		NewAccount().migration(m),
		NewSession().migration(m),
	})
}

func (m *modelModule) Injects() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//injections
		m.injectAccountTraits,
		m.injectSessionTraits,
	})
}
