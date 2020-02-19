package extend_traits

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
)

func (traits Traits) Migrate() error {
	return traits.Migrate_(traits.GormDB, traits.DormDB)
}

func (traits Traits) Migrate_(db *gorm.DB, dormDB *dorm.DB) error {
	err := db.AutoMigrate(traits.ObjectFactory()).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	model, err := dormDB.Model(traits.ObjectFactory().(dorm.ORMObject))
	if err != nil {
		return err
	}
	*traits.DormModel = *model
	return nil
}
