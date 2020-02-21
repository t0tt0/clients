package provider

import "github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"

func (s *Provider) TransactionDB() abstraction.TransactionDB {
	return s.transactionDB
}
