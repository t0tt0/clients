package vesclient

import "github.com/Myriad-Dreamin/go-ves/lib/fcg"

func (m modelModule) Migrates() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//migrations
		m.NewAccount().migration(m),
	})
}

func (m modelModule) Injects() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//injections
		m.injectAccountTraits,
	})
}
