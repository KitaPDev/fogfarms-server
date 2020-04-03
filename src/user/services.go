package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"github.com/golang/gddo/httputil/header"
	"golang.org/x/crypto/bcrypt"
)

var registeredUsers []models.User

func GetAllUsers() []models.User {
	if len(registeredUsers) == 0 {
		registeredUsers = repository.GetAllUsers()
	}

	return registeredUsers
}

func GetUser(username string) *models.User {
	if exists, user := Exists(username); exists {
		return user
	}

	return nil
}

func Exists(username string) (bool, *models.User) {
	for _, user := range registeredUsers {
		if user.Username == username {
			return true, &user
		}
	}
	return false, nil
}

func ValidateUserA(username string, password string) bool {
	return repository.ValidateUserA(username, password)
}

func hash(password string, salt string) string {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(h)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username string
		Password string
	}
	var testdata Input
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&testdata)

	fmt.Printf("%+v", testdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	repository.CreateUser(testdata.Username, testdata.Password)
}
