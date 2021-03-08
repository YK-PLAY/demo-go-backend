package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDbHelper struct {
	db *sql.DB
}

func (helper *MariaDbHelper) Init() {
	db, err := sql.Open("mysql", "test:test123@tcp(localhost:13306)/test")
	if err != nil {
		panic(err)
	}

	helper.db = db
}

func (helper *MariaDbHelper) Close() {
	if helper.db != nil {
		err := helper.db.Close()
		if err != nil {
			panic(err)
		}
	}
}

func (helper *MariaDbHelper) Insert(query string, args ...interface{}) int {
	r, err := helper.db.Exec(query, args)
	if err != nil {
		log.Fatalf("Insert error: %s\n", err.Error())
		return 0
	}

	n, err := r.RowsAffected()
	if err != nil {
		log.Fatalf("Insert affected error: %s\n", err.Error())
	}
	return int(n)
}

func (helper *MariaDbHelper) Select(query string, args ...interface{}) interface{} {
	rows, err := helper.db.Query(query)

	if err != nil {
		log.Fatalf("Select error: %s\n", err.Error())
		return nil
	}

	resultMap := make(map[int]interface{})
	for rows.Next() {
		var seq int64
		var idno string
		var regYmdt mysql.NullTime
		var modYmdt mysql.NullTime

		err := rows.Scan(&seq, &idno, &regYmdt, &modYmdt)
		if err != nil {
			log.Fatal(err)
		}

		m := map[string]interface{}{
			"seq":     seq,
			"idno":    idno,
			"regYmdt": regYmdt,
			"modYmdt": modYmdt,
		}

		index := len(resultMap)
		resultMap[index] = m
	}

	return resultMap
}
