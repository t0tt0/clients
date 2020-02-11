package vesclient

import "github.com/Myriad-Dreamin/go-ves/lib/fcg"

func (modelModule) Migrates() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//migrations
		Account{}.migrate,
	})
}

func (modelModule) Injects() error {
	return fcg.Calls([]fcg.MaybeInitializer{
		//injections
		injectAccountTraits,
	})
}
