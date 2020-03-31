package user

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

func GetUserFromRequest(r *http.Request) *models.User {
	username := r.Form.Get("username")

	return GetUser(username)
}

func Exists(username string) (bool, *models.User) {
	for _, user := range registeredUsers {
		if user.Username == username {
			return true, &user
		}
	}
	return false, nil
}

func ValidateUser(username string, password string) bool {
	if exists, user := Exists(username); exists {
		if user.Username == username {
			if user.Hash == hash(password, user.Salt) {
				return true
			}
		}
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