package model

import "database/sql"

type UserAuthLog struct {
	Seq       int64 `gorm:"primaryKey"`
	Cellphone string
	Number    string
	RegYmdt   sql.NullTime
}
