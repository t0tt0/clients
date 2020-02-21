package database

import "time"

type User struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
	LastLogin time.Time `dorm:"last_login" gorm:"column:last_login;default:CURRENT_TIMESTAMP;not null;" json:"last_login"`

	Name     string `dorm:"name" gorm:"column:name;unique;not_null" json:"name"`
	Password string `dorm:"password" gorm:"column:password;not_null" json:"password"`
}

// TableName specification
func (User) TableName() string {
	return "user"
}

func (d User) GetID() uint {
	return d.ID
}

func (d User) GetName() string {
	return d.Name
}

type UserFilter = Filter
