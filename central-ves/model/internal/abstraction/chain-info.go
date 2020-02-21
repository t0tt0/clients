package abstraction

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
)


// ChainInfoDB ...
type ChainInfoDB interface {
	GetTraits() ORMTraits
	Create(obj *database.ChainInfo) (aff int64, err error)
	Delete(obj *database.ChainInfo) (aff int64, err error)
	ID(id uint) (obj *database.ChainInfo, err error)
	Update(obj *database.ChainInfo) (aff int64, err error)

	UpdateFields(obj *database.ChainInfo, fields []string) (aff int64, err error)

	Query(opts ...ChainInfoQueryOption) (objs []database.ChainInfo, err error)
	Filter(f *database.ChainInfoFilter) (objs []database.ChainInfo, err error)
	FilterI(f interface{}) (obj interface{}, err error)

	InvertFind(acc uip.Account) (obj *database.ChainInfo, err error)
	FindAccounts(id uint, cid uip.ChainIDUnderlyingType) (as []uip.Account, err error)
	Scan(desc interface{}, opts ...ChainInfoQueryOption) (err error)
}

type ChainInfoQueryOption interface {
	implementsChainInfoQuery() ChainInfoQueryOption
}
