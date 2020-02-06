package extend_traits

import "github.com/jinzhu/gorm"

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

