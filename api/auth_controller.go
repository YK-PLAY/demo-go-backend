package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthControllerV1 struct {
	db *gorm.DB
}

func NewAuthControllerV1(db *gorm.DB) AuthControllerV1 {
	auth := AuthControllerV1{db: db}
	return auth
}

func (auth *AuthControllerV1) register(c *gin.Context) {
}
