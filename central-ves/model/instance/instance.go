package instance

import (
	"encoding/base64"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/errorc"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
)

type VESInstance struct {
	*model.Provider
}

func encodeBase64(que []byte) (res string) {
	return base64.StdEncoding.EncodeToString(que)
}

func (i VESInstance) InsertAccount(userName string, acc uiptypes.Account) error {
	// todo transaction
	userDB := i.UserDB()
	var user *model.User
	var err error
	if user, err = userDB.QueryName(userName); err != nil {
		return err
	} else if user == nil {
		user = &model.User{Name: userName}
		if err := errorc.CreateObj(user); err.Code != types.CodeOK {
			return err
		}
	}
	chainInfo := model.ChainInfo{
		UserID:  user.ID,
		Address: encodeBase64(acc.GetAddress()),
		ChainID: acc.GetChainId(),
	}
	if err := errorc.CreateObj(&chainInfo); err.Code != types.CodeOK {
		return err
	}
	return nil
}

func (i VESInstance) FindUser(userName string) (*model.User, error) {
	return i.UserDB().QueryName(userName)
}

func (i VESInstance) FindAccounts(userName string, chainID uint64) ([]uiptypes.Account, error) {
	user, err := i.UserDB().QueryName(userName)
	if err := errorc.MaybeSelectError(user, err); err.Code != types.CodeOK {
		return nil, err
	}

	return i.ChainInfoDB().FindAccounts(user.ID, chainID)
}

func (i VESInstance) HasAccount(userName string, acc uiptypes.Account) (has bool, err error) {
	var user *model.User
	user, err = i.InvertFind(acc)
	if err != nil {
		return false, err
	}
	if user == nil || user.Name != userName {
		return false, nil
	}
	return true, nil
}

func (i VESInstance) InvertFind(acc uiptypes.Account) (*model.User, error) {
	ci, err := i.ChainInfoDB().InvertFind(acc)
	if err != nil {
		return nil, err
	}
	if ci == nil {
		return nil, nil
	}
	return i.UserDB().ID(ci.ID)
}
