package abstraction

import (
	database2 "github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
)

type ObjectDB interface {
	GetTraits() ORMTraits
	Create(obj *database2.Object) (aff int64, err error)
	Delete(obj *database2.Object) (aff int64, err error)
	ID(id uint) (object *database2.Object, err error)
	Update(obj *database2.Object) (aff int64, err error)

	UpdateFields(obj *database2.Object, fields []string) (aff int64, err error)

	Query(opts ...ObjectQueryOption) (objs []database2.Object, err error)
	Filter(f *database2.ObjectFilter) (objs []database2.Object, err error)
	FilterI(f interface{}) (obj interface{}, err error)

	Scan(desc interface{}, opts ...ObjectQueryOption) (err error)
}

type ObjectQueryOption interface {
	implementsObjectQuery() ObjectQueryOption
}
