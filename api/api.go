package api

import (
	v1 "github.com/YK-PLAN/demo-go-backend/api/v1"
	"github.com/YK-PLAN/demo-go-backend/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	jwtMiddleware, err := jwt.New(&jwt.JwtMiddleware{})
	if err != nil {
		panic(err)
	}

	api := r.Group("/api")
	api.Use(jwtMiddleware.MiddlewareFunction())
	{
		v1.ApplyRoutes(api)
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, "")
		})
	}
}
