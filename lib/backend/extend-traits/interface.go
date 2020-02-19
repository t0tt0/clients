package extend_traits

import (
	"github.com/Myriad-Dreamin/dorm"
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
	"github.com/jinzhu/gorm"
)

type Traits struct {
	traits.ModelTraits
	ExtendModel
}

type Interface interface {
	traits.Interface
	Where2(template string) Where2Func
	Has1(template string) Has1Func
	Count1(template string) Count1Func
}

//ORMObject
type ORMObject = traits.ORMObject
type Where1Func = func(interface{}) (interface{}, error)
type Where1Func_ = func(db *gorm.DB, id interface{}) (interface{}, error)
type Where2Func = func(id, id2 interface{}) (interface{}, error)
type Has1Func = func(id interface{}) (has bool, err error)
type Has1Func_ = func(db *gorm.DB, id interface{}) (has bool, err error)

func NewTraits(t ORMObject, g *gorm.DB, d *dorm.DB) Traits {
	tt := traits.NewModelTraits(t, g, d)
	return Traits{
		ModelTraits: tt,
		ExtendModel: NewExtendModel(&tt),
	}
}
