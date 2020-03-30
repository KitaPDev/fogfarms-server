package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DbHost = "localhost"
	DbPort = 5432
	DbUser = "fogfarms"
	DbPass = "fogfarms"
	DbName = "fogfarms-01"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connectionString := fmt.Sprintf("port=%d host=%s user=%s " +
			"password=%s dbname=%s sslmode=disable",
			DbPort, DbHost, DbUser, DbPass, DbName)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			panic(err)
		}
	}

	return db
}