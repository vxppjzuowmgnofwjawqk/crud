package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	config := "user=postgres " +
		"password=130263 " +
		"host=localhost " +
		"port=5432 " +
		"dbname=crud " +
		"sslmode=disable"
	db, err := sql.Open("postgres", config)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
