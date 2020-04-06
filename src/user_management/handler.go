package user_management

import (
	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/user_management", populateUserManagementPage).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/user_management_register", register).
		Methods("POST").
		Schemes("http")

	return router
}

func register(w http.ResponseWriter, r *http.Request) {
	user.CreateUser(w, r)
}


func populateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	PopulateUserManagementPage(w, r)
}