package provider

import "github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"

func (s *Provider) TransactionDB() abstraction.TransactionDB {
	return s.transactionDB
}
