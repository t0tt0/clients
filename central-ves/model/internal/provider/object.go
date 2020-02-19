package provider

import "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"

func (s *Provider) ObjectDB() abstraction.ObjectDB {
	return s.objectDB
}
