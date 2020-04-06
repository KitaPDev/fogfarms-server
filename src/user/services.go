package user

import (
	"encoding/json"
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

var registeredUsers []models.User

func GetAllUsers() []models.User {
	if len(registeredUsers) == 0 {
		registeredUsers = repository.GetAllUsers()
	}

	return registeredUsers
}

func GetUserByUsername(username string) *models.User {
	if exists, user := ExistsByUsername(username); exists {
		return user
	}

	return nil
}
func GetUserByID(userID int) *models.User {
	if exists, user := ExistsByID(userID); exists {
		return user
	}

	return nil
}

func GetUsersByID(userIDs []int) []models.User {
	var users []models.User

	for _, userID := range userIDs {
		if exists, user := ExistsByID(userID); exists {
			users = append(users, *user)
		}
	}

	return users
}

func GetUserByUsernameFromRequest(w http.ResponseWriter, r *http.Request) *models.User {
	username := ""

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return nil
		}
	}
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		panic(err)
	}

	return GetUserByUsername(username)
}

func ExistsByUsername(username string) (bool, *models.User) {
	for _, user := range registeredUsers {
		if user.Username == username {
			return true, &user
		}
	}
	return false, nil
}

func ExistsByID(userID int) (bool, *models.User) {
	for _, user := range registeredUsers {
		if user.UserID == userID {
			return true, &user
		}
	}

	return false, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username string
		Password string
		IsAdministrator bool
	}

	var input Input
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&input)

	fmt.Printf("%+v", input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	repository.CreateUser(input.Username, input.Password, input.IsAdministrator)
}