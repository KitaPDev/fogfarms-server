package user

import (
	"encoding/json"
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func AuthenticateByUsername(username string, password string) (bool, error) {
	return repository.ValidateUserByUsername(username, password)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		IsAdministrator bool   `json:"is_administrator"`
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
	if err != nil {
		msg := "Error: Failed to Decode JSON"
		http.Error(w, msg, http.StatusBadRequest)
		log.Println(err)
	}

	err := repository.CreateUser(input.Username, input.Password, input.IsAdministrator)
	if err != nil {
		msg := "Error: Failed to Create User"
		http.Error(w, msg, http.StatusBadRequest)
		log.Println(err)
	}
}


func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(userID int) (*models.User, error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsersByID(userIDs []int) ([]models.User, error) {
	var users []models.User

	for _, userID := range userIDs {

		exists, user, err := ExistsByID(userID)
		if err != nil {
			return nil, err
		}

		if exists {
			users = append(users, *user)
		}
	}

	return users, nil
}

func GetUserByUsernameFromRequest(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	username := ""

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
		}
	}
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		msg := "Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
	}

	return GetUserByUsername(username)
}

func ExistsByUsername(username string) (bool, *models.User, error) {
	if user, err := GetUserByUsername(username); user != nil && err == nil {
		return true, user, nil
	} else {
		return false, nil, err
	}
}

func ExistsByID(userID int) (bool, *models.User, error) {
	if user, err := GetUserByID(userID); user != nil && err == nil {
		return true, user, nil
	} else {
		return false, nil, err
	}
}