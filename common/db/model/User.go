package model

import (
	"database/sql"
)

type User struct {
	Seq     int64 `gorm:"primaryKey"`
	Idno    string
	RegYmdt sql.NullTime
	ModYmdt sql.NullTime
}

func (User) TableName() string {
	return "users"
}
