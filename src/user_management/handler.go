package user_management

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/user"
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
	if !jwt.AuthenticateUser(w, r) {
		return
	}

	u := user.GetUserFromRequest(r)


}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}