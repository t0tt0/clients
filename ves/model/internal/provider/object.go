package provider

import "github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"

func (s *Provider) ObjectDB() abstraction.ObjectDB {
	return s.objectDB
}
