package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// const (
// 	DbHost = "localhost"
// 	DbPort = 5432
// 	DbUser = "fogfarms"
// 	DbPass = "fogfarms"
// 	DbName = "fogfarms-01"
// )
// const (
// 	DbHost = "localhost"
// 	DbPort = 5432
// 	DbUser = "postgres"
// 	DbPass = "postgres"
// 	DbName = "fogfarms-01"
// )

const (
	DbHost = "ec2-52-87-58-157.compute-1.amazonaws.com"
	DbPort = 5432
	DbUser = "ayvwvvempommus"
	DbPass = "6c88e08a05ae03e4b4a9c04de87e0933cce953a41d344d062390747b01cf673e"
	DbName = "dguhoh9d1tktp"
)

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connectionString := fmt.Sprintf("port=%d host=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			DbPort, DbHost, DbUser, DbPass, DbName)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			panic(err)
		}
	}

	return db
}
