package apiv1

import (
	"github.com/YK-PLAN/demo-go-backend/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
	}
}
