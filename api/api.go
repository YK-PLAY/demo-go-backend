package api

import (
	"github.com/YK-PLAN/demo-go-backend/middleware/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
	db  *gorm.DB
	jwt *jwt.JwtMiddleware
	v1  ApiV1
}

func NewApi(db *gorm.DB) Api {
	api := Api{db: db}
	return api
}

func (api *Api) ApplyRoutes(r *gin.Engine) {
	jwtMiddleware, err := jwt.New(&jwt.JwtMiddleware{})
	api.jwt = jwtMiddleware
	if err != nil {
		panic(err)
	}

	apiGroup := r.Group("/api")
	apiGroup.Use(jwtMiddleware.MiddlewareFunction())
	{
		v1 := ApiV1{db: api.db}
		api.v1 = v1
		v1.ApplyRoutes(apiGroup)
	}
}
