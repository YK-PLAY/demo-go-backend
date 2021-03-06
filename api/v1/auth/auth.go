package auth

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/register/auth", registerAuth)
		auth.POST("/parse", parseToken)
	}
}
