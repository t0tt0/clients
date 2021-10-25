package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func wrapToObject(object interface{}, err error) (*database.Object, error) {
	if object == nil {
		return nil, err
	}
	return object.(*database.Object), err
}

type ObjectDB struct {
	traits abstraction.ORMTraits
}

func NewObjectDB(cb func(interface{}) abstraction.ORMTraits, _ module.Module) (*ObjectDB, error) {
	return &ObjectDB{traits: cb(Object{})}, nil
}

func (objectDB ObjectDB) GetTraits() abstraction.ORMTraits {
	return objectDB.traits
}

func (objectDB ObjectDB) Query(options ...abstraction.ObjectQueryOption) (objs []database.Object, err error) {
	panic("implement me")
}

func (objectDB ObjectDB) Scan(desc interface{}, options ...abstraction.ObjectQueryOption) (err error) {
	panic("implement me")
}

func (objectDB ObjectDB) Migrate() error {
	return objectDB.traits.Migrate()
}

func (objectDB ObjectDB) Create(d *database.Object) (aff int64, err error) {
	return objectDB.traits.Create(d)
}

func (objectDB ObjectDB) Update(d *database.Object) (aff int64, err error) {
	return objectDB.traits.Update(d)
}

func (objectDB ObjectDB) UpdateFields(d *database.Object, fields []string) (aff int64, err error) {
	return objectDB.traits.UpdateFields(d, fields)
}

func (objectDB ObjectDB) UpdateFields_(d *database.Object, db *dorm.DB, fields []string) (aff int64, err error) {
	return objectDB.traits.UpdateFields_(db, d, fields)
}

func (objectDB ObjectDB) UpdateFields__(d *database.Object, db dorm.SQLCommon, fields []string) (aff int64, err error) {
	return objectDB.traits.UpdateFields__(db, d, fields)
}

func (objectDB ObjectDB) Delete(d *database.Object) (aff int64, err error) {
	return objectDB.traits.Delete(d)
}

func (objectDB *ObjectDB) Filter(f *database.Filter) (user []database.Object, err error) {
	err = objectDB.traits.Filter(f, &user)
	return
}

func (objectDB *ObjectDB) FilterI(f interface{}) (interface{}, error) {
	return objectDB.Filter(f.(*database.Filter))
}

func (objectDB *ObjectDB) ID(id uint) (object *database.Object, err error) {
	return wrapToObject(objectDB.traits.ID(id))
}

func (objectDB *ObjectDB) ID_(db *gorm.DB, id uint) (object *database.Object, err error) {
	return wrapToObject(objectDB.traits.ID_(db, id))
}

type ObjectQuery struct {
	db *gorm.DB
}


func (objectDB *ObjectDB) QueryChain() *ObjectQuery {
	return &ObjectQuery{db: objectDB.traits.GetGormDB()}
}

func (objectDB *ObjectQuery) Order(order string, reorder ...bool) *ObjectQuery {
	objectDB.db = objectDB.db.Order(order, reorder...)
	return objectDB
}

func (objectDB *ObjectQuery) Page(page, pageSize int) *ObjectQuery {
	objectDB.db = objectDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return objectDB
}
func (objectDB *ObjectQuery) BeforeID(id uint) *ObjectQuery {
	objectDB.db = objectDB.db.Where("id <= ?", id)
	return objectDB
}

func (objectDB *ObjectQuery) Preload() *ObjectQuery {
	objectDB.db = objectDB.db.Set("gorm:auto_preload", true)
	return objectDB
}

func (objectDB *ObjectQuery) Query() (objects []database.Object, err error) {
	err = objectDB.db.Find(&objects).Error
	return
}

func (objectDB *ObjectQuery) Scan(desc interface{}) (err error) {
	err = objectDB.db.Scan(desc).Error
	return
}
