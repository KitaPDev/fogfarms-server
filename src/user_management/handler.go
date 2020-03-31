package user_management

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/user_management", getAllUsers).
		Methods("GET").
		Schemes("http")

	return router
}

func populateUserManagementPage(w http.ResponseWriter, r *http.Request) {

}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}