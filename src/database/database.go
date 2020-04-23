package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var db *sql.DB

//const (
//	dbHost = "localhost"
//	dbPort = 5432
//	dbUser = "fogfarms"
//	dbPass = "fogfarms"
//	dbName = "fogfarms-01"
//	sslMode = "disable"
//)

//const (
//	DbHost  = "localhost"
//	DbPort  = 5432
//	DbUser  = "postgres"
//	DbPass  = "postgres"
//	DbName  = "fogfarms-01"
//	SSLMODE = "disable"
//)

func GetDB() *sql.DB {
	var err error

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		return nil
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	if db == nil {
		connectionString := fmt.Sprintf("port=%d host=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			dbPort, dbHost, dbUser, dbPass, dbName, sslMode)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			panic(err)
		}
	}

	return db
}
