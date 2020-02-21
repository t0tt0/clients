package abstraction

import "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"

type UserDB interface {
	GetTraits() ORMTraits
	Create(u *database.User) (aff int64, err error)
	Update(u *database.User) (aff int64, err error)
	ID(id uint) (u *database.User, err error)
	Delete(u *database.User) (aff int64, err error)

	UpdateFields(u *database.User, fields []string) (aff int64, err error)

	Register(u *database.User) (aff int64, err error)
	ResetPassword(u *database.User, pwd string) (aff int64, err error)
	AuthenticatePassword(u *database.User, pwd string) (authenticated bool, err error)

	Has(id uint) (has bool, err error)
	HasName(id string) (has bool, err error)
	QueryName(id string) (u *database.User, err error)

	Query(opts ...UserQueryOption) (u *database.User, err error)
	Filter(f *database.Filter) (u []database.User, err error)
	FilterI(f interface{}) (u interface{}, err error)
	Scan(desc interface{}, opts ...UserQueryOption) (err error)
}

type UserQueryOption interface {
	implementsUserQuery() UserQueryOption
}
