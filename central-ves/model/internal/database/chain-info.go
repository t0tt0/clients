package database

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"time"
)

type ChainInfo struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	UserID  uint                           `dorm:"user_id" gorm:"column:user_id;not_null" json:"user_id"`
	Address string                         `dorm:"address" gorm:"address;not_null" json:"address"`
	ChainID uip.ChainIDUnderlyingType `dorm:"chain_id" gorm:"chain_id;not_null" json:"chain_id"`
}

// TableName specification
func (ChainInfo) TableName() string {
	return "chain_info"
}

func (ci ChainInfo) GetID() uint {
	return ci.ID
}

type ChainInfoFilter = Filter
