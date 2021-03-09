package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/YK-PLAN/demo-go-backend/common/db/model"
	"gorm.io/gorm"
)

func GetUserAuthLogByCellphone(db *gorm.DB, cellphone string) model.UserAuthLog {
	var log model.UserAuthLog
	checkTime := time.Now().Local().Add(-3 * time.Minute)
	db.Where(model.UserAuthLog{Cellphone: cellphone, RegYmdt: sql.NullTime{Time: checkTime, Valid: true}}).Last(&log)
	return log
}

func SaveUserAuthLog(db *gorm.DB, cellphone string, number string) int {
	authLog := model.UserAuthLog{
		Cellphone: cellphone,
		Number:    number,
		RegYmdt: sql.NullTime{
			Time:  time.Now().Local(),
			Valid: true,
		},
	}

	tx := db.Save(&authLog)
	if tx.Error != nil {
		log.Printf("SaveUserAuthLog failed: %s\n", tx.Error.Error())
		return 0
	}

	return int(tx.RowsAffected)
}
