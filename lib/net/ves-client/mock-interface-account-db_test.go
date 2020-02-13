package vesclient

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
)

type mockAccountDBImpl struct {
	queryAliasMap map[string]*Account
}

func (m mockAccountDBImpl) Create(ci *Account) (int64, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) Update(ci *Account) (int64, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) UpdateFields(ci *Account, fields []string) (int64, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) UpdateFields_(ci *Account, db *dorm.DB, fields []string) (int64, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) UpdateFields__(ci *Account, db dorm.SQLCommon, fields []string) (int64, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) Delete(ci *Account) (int64, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) ID(id uint) (account *Account, err error) {
	panic("implement me")
}

func (m mockAccountDBImpl) ID_(db *gorm.DB, id uint) (account *Account, err error) {
	panic("implement me")
}

func (m mockAccountDBImpl) InvertFind(acc uiptypes.Account) (account *Account, err error) {
	panic("implement me")
}

func (m mockAccountDBImpl) QueryAlias(alias string) (account *Account, err error) {
	if account, ok := m.queryAliasMap[alias]; ok {
		return account, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (m mockAccountDBImpl) FindAccounts(id uint, chainID uiptypes.ChainIDUnderlyingType) ([]uiptypes.Account, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) QueryChain() *AccountQuery {
	panic("implement me")
}

type AccountMockData interface{}

type AccountQueryAliasMockData struct {
	K string
	V *Account
}

func mockAccountDB(data ...AccountMockData) AccountDBInterface {
	accountDB := mockAccountDBImpl{
		queryAliasMap:make(map[string]*Account),
	}
	for i := range data {
		switch d := data[i].(type) {
		case AccountQueryAliasMockData:
			accountDB.queryAliasMap[d.K] = d.V
		}
	}
	return accountDB
}

