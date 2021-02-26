package api

import (
	v1 "github.com/YK-PLAN/demo-go-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
