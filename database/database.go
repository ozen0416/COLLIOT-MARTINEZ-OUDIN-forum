package database

import (
	"database/sql"
	"fmt"
)

func ReturnDatabase() *sql.DB {
	// Connecting to BD
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
