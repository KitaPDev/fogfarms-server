package user_management

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/user_management", populateUserManagementPage).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/user_management/register", register).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/user_management/assign/", assignUserModuleGroupPermission).
		Methods("POST").
		Schemes("http")

	return router
}
// todo ddfsd - test
func register(w http.ResponseWriter, r *http.Request) {
	CreateUser(w, r)
}

// todo ddfsd - test
func populateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	PopulateUserManagementPage(w, r)
}

// todo ddfsd - test
func assignUserModuleGroupPermission(w http.ResponseWriter, r *http.Request) {
	AssignUserModuleGroupPermission(w, r)
}