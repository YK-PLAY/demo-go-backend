package main

import (
	"flag"

	"github.com/YK-PLAN/demo-go-backend/api"
	"github.com/YK-PLAN/demo-go-backend/common/db"
	"github.com/YK-PLAN/demo-go-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	env := flag.String("env", "local", "Current environmnet")
	flag.Parse()

	cnf := config.Load(*env)
	run(&cnf)
}

func run(cnf *config.Config) {
	r := gin.Default()

	db := db.NewMysqlDB()
	api := api.NewApi(db)
	api.ApplyRoutes(r)

	r.Run(cnf.Address())
}
