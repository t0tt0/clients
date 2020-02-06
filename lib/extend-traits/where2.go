package extend_traits

import "github.com/jinzhu/gorm"

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

