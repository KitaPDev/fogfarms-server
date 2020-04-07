package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.User, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

	rows, err := db.Query("SELECT * FROM Users WHERE Username = $1;", username)
	if err != nil {
		return nil, err
	}
  
	defer rows.Close()

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
	defer rows.Close()

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

func ValidateUserByUsername(username string, inputPassword string) (bool, error) {
	db := database.GetDB()

	sqlStatement := `SELECT UserID, Username, Hash, Salt FROM Users WHERE Username = $1;`

	user := models.User{}

	row := db.QueryRow(sqlStatement, username)

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
			return true, nil
		}
		return false, nil

	default:
		fmt.Printf("%+v", err)
		return false, nil
	}

	return false, nil
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
