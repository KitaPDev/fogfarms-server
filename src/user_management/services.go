package user_management

import (
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	repository.GetAllUsers()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func PopulateUserManagementPage(w http.ResponseWriter, r *http.Request) {

}