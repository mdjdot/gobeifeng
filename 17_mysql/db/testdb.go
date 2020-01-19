package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

// TestDB .
var TestDB *sql.DB
var err error

func init() {
	TestDB, err = sql.Open("mysql", "root:dmtest@tcp(localhost:3306)/testdb?charset=utf8")
	if err != nil {
		panic(err)
	}
}
