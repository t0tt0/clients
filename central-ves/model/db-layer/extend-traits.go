package dblayer

import (
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
	"github.com/jinzhu/gorm"
)

type ExtendModelOperationInterface = traits.Interface

type ExtendModel struct {
	i           ExtendModelOperationInterface
	replacement interface{}
}

func NewExtendModel(t ExtendModelOperationInterface) ExtendModel {
	return ExtendModel{
		i:           t,
		replacement: t.ObjectFactory(),
	}
}

type Where2Func = func(id, id2 interface{}) (interface{}, error)

func (model ExtendModel) Where2(template string) Where2Func {
	return func(id, id2 interface{}) (object interface{}, err error) {
		object = model.i.ObjectFactory()
		rdb := model.i.GetGormDB().Where(template, id, id2).Find(object)
		err = rdb.Error
		if rdb.RecordNotFound() {
			object = nil
			err = nil
		}
		return
	}
}

func (model ExtendModel) Where2_(template string) func(db *gorm.DB, id, id2 interface{}) (interface{}, error) {
	return func(db *gorm.DB, id, id2 interface{}) (object interface{}, err error) {
		object = model.i.ObjectFactory()
		rdb := db.Where(template, id, id2).Find(object)
		err = rdb.Error
		if rdb.RecordNotFound() {
			object = nil
			err = nil
		}
		return
	}
}

type Has1Func = func(id interface{}) (has bool, err error)

func (model ExtendModel) Has1(template string) Has1Func {
	whereF := model.i.Where1(template)
	return func(id interface{}) (has bool, err error) {
		obj, err := whereF(id)
		if err != nil {
			return false, err
		} else if obj == nil {
			return false, nil
		} else {
			return true, nil
		}
	}
}

type Has1Func_ = func(db *gorm.DB, id interface{}) (has bool, err error)

func (model ExtendModel) Has1_(template string) Has1Func_ {
	whereF := model.i.Where1_(template)
	return func(db *gorm.DB, id interface{}) (has bool, err error) {
		obj, err := whereF(db, id)
		if err != nil {
			return false, err
		} else if obj == nil {
			return false, nil
		} else {
			return true, nil
		}
	}
}
