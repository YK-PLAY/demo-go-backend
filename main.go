package main

import (
	"flag"
	"log"

	"github.com/YK-PLAN/demo-go-backend/api"
	"github.com/YK-PLAN/demo-go-backend/common/db"
	"github.com/YK-PLAN/demo-go-backend/common/db/repository"
	"github.com/YK-PLAN/demo-go-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	env := flag.String("env", "local", "Current environmnet")
	flag.Parse()

	cnf := config.Load(*env)

	var helper db.MariaDbHelper
	helper.Init()
	defer helper.Close()

	r := helper.Select("Select * from users", nil)
	log.Printf("Select result: %+v\n", r)

	run(&cnf)
}

func run(cnf *config.Config) {
	r := gin.Default()

	db := db.NewMysqlDB()
	a := api.NewApi(db)
	a.ApplyRoutes(r)

	user := repository.GetUserBySeq(db, 1)
	log.Printf("User: %+v\n", user)

	r.Run(cnf.Address())
}
