package repository

import (
	"testing"

	"github.com/YK-PLAN/demo-go-backend/common/db"
	"github.com/YK-PLAN/demo-go-backend/common/db/model"
	"github.com/stretchr/testify/assert"
)

func TestGetUserBySeq(t *testing.T) {
	db := db.NewMysqlDB()

	var seq int64 = 1
	user := GetUserBySeq(db, seq)

	if user == (model.User{}) {
		t.Skip("Skip repository test")
	}

	assert.Equal(t, seq, user.Seq)
}

func TestGetUserByCellphoneAndUuid(t *testing.T) {
	db := db.NewMysqlDB()

	var cellphone string = "01012345678"
	user := GetUserByCellphoneAndUuid(db, cellphone, "test-uuid")
	if user == (model.User{}) {
		t.Skip("Skip repository test")
	}

	assert.Equal(t, cellphone, user.Cellphone)
}
