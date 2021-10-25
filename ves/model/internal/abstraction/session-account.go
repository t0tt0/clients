package abstraction

import (
	database2 "github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
)

// SessionAccountDB ...
type SessionAccountDB interface {
	GetTraits() ORMTraits

	Create(sa *database2.SessionAccount) (aff int64, err error)
	Delete(sa *database2.SessionAccount) (aff int64, err error)
	Find(sa *database2.SessionAccount) (has bool, err error)
	UpdateAcknowledged(sa *database2.SessionAccount) (aff int64, err error)

	GetAcknowledged(guid string) (aff int64, err error)
	GetTotal(guid string) (aff int64, err error)
	ID(id string) (sas []database2.SessionAccount, err error)

	Filter(f *database2.SessionAccountFilter) (sas []database2.SessionAccount, err error)
	FilterI(f interface{}) (interface{}, error)

	Query(opts ...SessionAccountQueryOption) (sas []database2.SessionAccount, err error)
	Scan(obj interface{}, opts ...SessionAccountQueryOption) (err error)
}

type SessionAccountQueryOption interface {
	implementsSessionAccountQuery() SessionAccountQueryOption
}
