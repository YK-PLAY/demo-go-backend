package repository

import (
	"github.com/YK-PLAN/demo-go-backend/common/db/model"
	"gorm.io/gorm"
)

func GetUserBySeq(db *gorm.DB, seq int64) model.User {
	var user model.User
	db.Where(model.User{Seq: 1}).Find(&user)
	return user
}
