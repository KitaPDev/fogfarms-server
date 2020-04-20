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
// 	SSLMODE = "disable"
// )
const (
	DbHost  = "localhost"
	DbPort  = 5432
	DbUser  = "postgres"
	DbPass  = "postgres"
	DbName  = "fogfarms-01"
	SSLMODE = "disable"
)

//const (
//	DbHost = "ec2-54-157-78-113.compute-1.amazonaws.com"
//	DbPort = 5432
//	DbUser = "aevojjwxgydmym"
//	DbPass = "f3940dee95e22cb02e4dea372a8252fa356265e2b3fde6875307ed563e88f639"
//	DbName = "dbfi6i4j26c9jj"
//	SSLMODE = "require"
//)

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connectionString := fmt.Sprintf("port=%d host=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			DbPort, DbHost, DbUser, DbPass, DbName, SSLMODE)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			panic(err)
		}
	}

	return db
}
