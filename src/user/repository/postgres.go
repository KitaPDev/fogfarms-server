package repository

import (
	"database/sql"
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetAllUsers() []models.User {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		panic(err)
	}
	defer log.Fatal(rows.Close())

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Hash,
			&user.Salt,
			&user.IsAdministrator,
			&user.CreatedAt,
		)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func CreateUser(username string, password string) {
	db := database.GetDB()
	sqlStatement := fmt.Sprintf("INSERT INTO Users (Username, IsAdministrator, Hash, Salt, CreatedAt)" +
		"VALUES ($1, False , $2, 's', Now())\n" +
		"RETURNING Username, Hash;")
	Username := ""
	Hash := ""
	hashInset := hash(password, "s")
	err := db.QueryRow(sqlStatement, username, hashInset).Scan(&Username, &Hash)
	if err != nil {
		panic(err)
	}
	fmt.Println(Username, "\n hash:", Hash)
}

func ValidateUserA(usernameIn string, password string) bool {
	db := database.GetDB()
	sqlStatement := `SELECT username , hash,salt FROM users WHERE username=$1;`
	//sqlStatement := `SELECT username , hash,salt FROM users;`
	var username string
	var salt string
	var hash string
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.

	row := db.QueryRow(sqlStatement, usernameIn)
	switch err := row.Scan(&username, &hash, &salt); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		actualpassword := password + salt
		fmt.Println(username, salt)
		fmt.Printf("this works \n")
		fmt.Printf("%+v , %+v \n", username, actualpassword)
		if usernameIn == username && bcrypt.CompareHashAndPassword([]byte(hash), []byte(actualpassword)) == nil {
			return true
		}
		return false
	default:
		panic(err)
	}

	return false
}

func hash(password string, salt string) string {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(h)
}
