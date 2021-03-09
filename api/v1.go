package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiV1 struct {
	db   *gorm.DB
	auth AuthControllerV1
}

func (v1 *ApiV1) ApplyRoutes(r *gin.RouterGroup) {
	v1Group := r.Group("/v1")
	{
		auth := NewAuthControllerV1(v1.db)
		v1.auth = auth
		v1Group.POST("/auth/register", auth.register)
	}
}
