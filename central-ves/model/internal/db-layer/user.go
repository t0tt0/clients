package dblayer

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
	"github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func wrapToUser(user interface{}, err error) (*User, error) {
	if user == nil {
		return nil, err
	}
	return user.(*User), err
}

type UserDB struct {
	traits abstraction.ORMTraits

	queryNameFunc extend_traits.Where1Func
	hasNameFunc   extend_traits.Has1Func
}

func (userDB UserDB) GetTraits() abstraction.ORMTraits {
	return userDB.traits
}

func (userDB UserDB) Query(opts ...abstraction.UserQueryOption) (u *database.User, err error) {
	panic("implement me")
}

func (userDB UserDB) Scan(desc interface{}, opts ...abstraction.UserQueryOption) (err error) {
	panic("implement me")
}

func NewUserDB(cb func(interface{}) abstraction.ORMTraits, _ module.Module) (*UserDB, error) {
	traits := cb(User{})
	return &UserDB{
		traits:        traits,
		queryNameFunc: traits.Where1("name = ?"),
		hasNameFunc:   traits.Has1("name = ?"),
	}, nil
}

func (userDB UserDB) Migrate() error {
	return userDB.traits.Migrate()
}

func (userDB UserDB) Create(d *User) (aff int64, err error) {
	return userDB.traits.Create(d)
}

func (userDB UserDB) Update(d *User) (aff int64, err error) {
	return userDB.traits.Update(d)
}

func (userDB UserDB) UpdateFields(d *User, fields []string) (aff int64, err error) {
	return userDB.traits.UpdateFields(d, fields)
}

func (userDB UserDB) Delete(d *User) (aff int64, err error) {
	return userDB.traits.Delete(d)
}

func (userDB UserDB) Register(d *User) (aff int64, err error) {
	d.Password, err = crypto.NewPasswordString(d.Password)
	if err != nil {
		return 0, err
	}

	return userDB.traits.Create(d)
}

func (userDB UserDB) ResetPassword(d *User, password string) (aff int64, err error) {
	d.Password, err = crypto.NewPasswordString(password)
	if err != nil {
		return 0, err
	}

	return userDB.traits.UpdateFields(d, []string{"password"})
}

func (userDB UserDB) AuthenticatePassword(d *User, pswd string) (bool, error) {
	return crypto.CheckPasswordString(pswd, d.Password)
}

func (userDB *UserDB) Filter(f *database.Filter) (user []User, err error) {
	err = userDB.traits.Filter(f, &user)
	return
}

func (userDB *UserDB) FilterI(f interface{}) (interface{}, error) {
	return userDB.Filter(f.(*database.Filter))
}

func (userDB *UserDB) ID(id uint) (user *User, err error) {
	return wrapToUser(userDB.traits.ID(id))
}

type UserQuery struct {
	db *gorm.DB
}

func (userDB *UserDB) QueryChain() *UserQuery {
	return &UserQuery{db: userDB.traits.GetGormDB()}
}

func (userDB *UserQuery) Order(order string, reorder ...bool) *UserQuery {
	userDB.db = userDB.db.Order(order, reorder...)
	return userDB
}

func (userDB *UserQuery) Page(page, pageSize int) *UserQuery {
	userDB.db = userDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return userDB
}
func (userDB *UserQuery) BeforeID(id uint) *UserQuery {
	userDB.db = userDB.db.Where("id <= ?", id)
	return userDB
}

func (userDB *UserQuery) Preload() *UserQuery {
	userDB.db = userDB.db.Set("gorm:auto_preload", true)
	return userDB
}

func (userDB *UserQuery) Query() (users []User, err error) {
	err = userDB.db.Find(&users).Error
	return
}

func (userDB *UserDB) Has(id uint) (has bool, err error) {
	return userDB.traits.Has(id)
}

func (userDB *UserDB) HasName(id string) (has bool, err error) {
	return userDB.hasNameFunc(id)
}

func (userDB *UserDB) QueryName(id string) (user *User, err error) {
	return wrapToUser(userDB.queryNameFunc(id))
}
