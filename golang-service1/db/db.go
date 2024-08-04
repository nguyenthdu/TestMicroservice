package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=password dbname=servicetest_db sslmode=disable")
	if err != nil {
		panic(err)
	}
}
