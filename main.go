package main

import (
	"github.com/YK-PLAN/demo-go-backend/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api.ApplyRoutes(r)

	r.Run()
}
