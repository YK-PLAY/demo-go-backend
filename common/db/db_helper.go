package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jmrobles/h2go"
)

type dbHelper interface {
	Init()
	Insert(tablename string, entity Entity) int
	Select() interface{}
}

var (
	helper dbHelper
)

func InitDBHelper(env string) {
	if env == "local" {
		helper = h2DbHelper{}
	} else {
		fmt.Printf("Not supported env: %s\n", env)
	}
}

func Inert(tablename string, entity Entity) int {
	conn, err := sql.Open("h2", "h2://sa@localhost/test?logging=info")
	if err != nil {
		log.Fatalf("Connect to database error: %s", err)
	}

	// Create table
	log.Printf("CREATE TABLE")
	ret, err := conn.Exec("CREATE TABLE test (id int not null, name varchar(100))")
	if err != nil {
		log.Printf("Can't execute sentence: %s", err)
		return 0
	}

	log.Printf("%v\n", ret)

	// Insert
	ret, err = conn.Exec("INSERT INTO test VALUES (?, ?)",
		1, "John")
	if err != nil {
		log.Printf("Can't execute sentence: %s", err)
		return 0
	}

	return helper.Insert(tablename, entity)
}
