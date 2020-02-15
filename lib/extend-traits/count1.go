package extend_traits

import "github.com/jinzhu/gorm"

type Count1Func = func(id interface{}) (count int64, err error)
type Count1Func_ = func(db *gorm.DB, id interface{}) (count int64, err error)

func (model ExtendModel) Count1(template string) Count1Func {
	return func(id interface{}) (count int64, err error) {
		rdb := model.i.GetGormDB().Model(model.replacement).Where(template, id).Count(&count)
		err = rdb.Error
		return
	}
}

func (model ExtendModel) Count1_(template string) Count1Func_ {
	return func(db *gorm.DB, id interface{}) (count int64, err error) {
		rdb := db.Where(template, id).Model(model.replacement).Count(&count)
		err = rdb.Error
		return
	}
}
