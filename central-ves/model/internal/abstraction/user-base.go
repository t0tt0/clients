package abstraction

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
)

// the database which used by others
type UserBase interface {
	// insert accounts maps from guid to account
	InsertAccount(userName string, acc uip.Account) error

	// DeleteAccount(userName string, Account) error

	// DeleteAccountByName(userName string) error

	// DeleteAccountByID(user_id) error

	// find accounts which guid is corresponding to user
	FindUser(userName string) (*database.User, error)

	// find accounts which guid is corresponding to user
	FindAccounts(userName string, chainID uint64) ([]uip.Account, error)

	// return true if user has this account
	HasAccount(userName string, acc uip.Account) (has bool, err error)

	// return the user which has this account
	InvertFind(uip.Account) (*database.User, error)
}
