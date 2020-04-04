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

func CreateUser(username string, password string, isAdministrator bool) {
	db := database.GetDB()

	salt := generateSalt()
	hash := hash(password, salt)

	sqlStatement := fmt.Sprintf("INSERT INTO Users (Username, IsAdministrator, Hash, Salt, CreatedAt)" +
		"VALUES (%s, %v, %s, %s, Now())",
		username, isAdministrator, hash, salt)

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func ValidateUser(username string, inputPassword string) bool {
	db := database.GetDB()

	sqlStatement := fmt.Sprintf("SELECT UserID, Username, Hash, Salt FROM Users WHERE Username = %s;",
		username)

	user := models.User{}

	row, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	switch err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Hash,
		&user.Salt,
		); err {

		case sql.ErrNoRows:
			fmt.Println("No Rows Returned!")

		case nil:
			password := inputPassword + user.Salt
			fmt.Println(username, user.Salt)
			fmt.Printf("this works \n")
			fmt.Printf("%+v , %+v \n", username, password)

			if username == username &&
				bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password)) == nil {
				return true
			}
			return false

		default:
			panic(err)
	}

	return false
}

func generateSalt() string {
	return string(make([]byte, 32))
}

func hash(password string, salt string) string {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(h)
}
