package vesclient

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/lib/basic/encoding"
	"github.com/jinzhu/gorm"
)

type mockAccountDBImpl struct {
	queryAliasMap map[string]*Account
	invertFindMap map[string]*Account
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

func (m mockAccountDBImpl) InvertFind(acc uip.Account) (account *Account, err error) {
	if account, ok := m.invertFindMap[m.invertFindKey(acc)]; ok {
		return account, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (m mockAccountDBImpl) QueryAlias(alias string) (account *Account, err error) {
	if account, ok := m.queryAliasMap[alias]; ok {
		return account, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (m mockAccountDBImpl) FindAccounts(id uint, chainID uip.ChainIDUnderlyingType) ([]uip.Account, error) {
	panic("implement me")
}

func (m mockAccountDBImpl) QueryChain() *AccountQuery {
	panic("implement me")
}

func (m mockAccountDBImpl) invertFindKey(account uip.Account) string {
	const mg = "<84f4446f>"
	hashTarget := make([]byte, len(mg)+len(account.GetAddress())+8)
	binary.BigEndian.PutUint64(hashTarget, account.GetChainId())
	copy(hashTarget[8:], mg)
	copy(hashTarget[18:], account.GetAddress())
	hash := md5.Sum(hashTarget)
	return encoding.EncodeHex(hash[:])
}

type AccountMockData interface{}

type AccountQueryAliasMockData struct {
	K string
	V *Account
}

type accountKey = uip.AccountImpl
type AccountInvertFindMockData struct {
	K *accountKey
	V *Account
}

func mockAccountDB(data ...AccountMockData) AccountDBInterface {
	accountDB := mockAccountDBImpl{
		queryAliasMap: make(map[string]*Account),
		invertFindMap: make(map[string]*Account),
	}
	for i := range data {
		switch d := data[i].(type) {
		case AccountQueryAliasMockData:
			accountDB.queryAliasMap[d.K] = d.V
		case AccountInvertFindMockData:
			accountDB.invertFindMap[accountDB.invertFindKey(d.K)] = d.V
		}
	}
	return accountDB
}
