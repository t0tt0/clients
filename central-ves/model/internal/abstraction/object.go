package abstraction

import "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"

type ObjectDB interface {
	GetTraits() ORMTraits
	Create(obj *database.Object) (aff int64, err error)
	Delete(obj *database.Object) (aff int64, err error)
	ID(id uint) (object *database.Object, err error)
	Update(obj *database.Object) (aff int64, err error)

	UpdateFields(obj *database.Object, fields []string) (aff int64, err error)

	Query(opts ...ObjectQueryOption) (objs []database.Object, err error)
	Filter(f *database.ObjectFilter) (objs []database.Object, err error)
	FilterI(f interface{}) (obj interface{}, err error)

	Scan(desc interface{}, opts ...ObjectQueryOption) (err error)
}

type ObjectQueryOption interface {
	implementsObjectQuery() ObjectQueryOption
}
