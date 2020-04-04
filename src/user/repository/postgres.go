package repository

import (
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

func hash(password string, salt string) string {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(h)
}