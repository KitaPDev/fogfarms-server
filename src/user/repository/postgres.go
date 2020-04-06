package repository

import (
	"database/sql"
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

func GetAllUsers() ([]models.User, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		return nil, err
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
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Users WHERE Username = ?;", username)
	if err != nil {
		return nil, err
	}
	defer log.Fatal(rows.Close())

	var user models.User
	for rows.Next() {
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Hash,
			&user.Salt,
			&user.IsAdministrator,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

	}

	return &user, nil
}

func GetUserByID(userID int) (*models.User, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Users WHERE UserID = ?;", userID)
	if err != nil {
		return nil, err
	}
	defer log.Fatal(rows.Close())

	var user models.User
	for rows.Next() {
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Hash,
			&user.Salt,
			&user.IsAdministrator,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

	}

	return &user, nil
}

func CreateUser(username string, password string, isAdministrator bool) {
	db := database.GetDB()

	salt := generateSalt()
	hash := hash(password, salt)

	sqlStatement := `INSERT INTO Users (Username, IsAdministrator, Hash, Salt, CreatedAt) 
		VALUES ($1, $2, $3, $4, Now());`

	db.QueryRow(sqlStatement, username, isAdministrator, hash, salt)
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
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	salt := make([]byte, 32)
	for i := range salt {
		salt[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(salt)
}

func hash(password string, salt string) string {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(h)
}
