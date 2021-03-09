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

func GetUserByCellphone(db *gorm.DB, cellphone string) model.User {
	var user model.User
	db.Where(model.User{Cellphone: cellphone}).Find(&user)
	return user
}

func GetUserByCellphoneAndUuid(db *gorm.DB, cellphone string, uuid string) model.User {
	var user model.User
	db.
		Where(model.User{Cellphone: cellphone}).
		Joins("INNER JOIN user_uuid ON ? = user_uuid.uuid AND users.seq = user_uuid.user_seq", uuid).
		First(&user)
	return user
}
