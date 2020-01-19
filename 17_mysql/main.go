package main

import (
	"fmt"

	"github.com/mbtdot/gobeifeng/17_mysql/db"
)

func main() {
	var id int
	var name string

	stmt, err := db.TestDB.Prepare("select * from user_info")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
