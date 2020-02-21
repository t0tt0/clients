package database

import "time"

type Transaction struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	SessionID string `dorm:"session_id" gorm:"column:session_id;not null;" json:"session_id"`
	Index     int64  `dorm:"idx" gorm:"column:idx;default:0;not null;" json:"idx"`
	Content   string `dorm:"contents" gorm:"column:content;type:text;not null;" json:"content"`
}

// TableName specification
func (Transaction) TableName() string {
	return "transaction"
}

func (d Transaction) GetID() uint {
	return d.ID
}

type TransactionFilter = Filter
