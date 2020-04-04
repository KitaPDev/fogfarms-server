package user_management

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/user_management", populateUserManagementPage).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/user_management_register", user.CreateUser).
		Methods("POST").
		Schemes("http")

	return router
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func populateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUser(w, r) {
		return
	}

	u := user.GetUserFromRequest(r)

	if u.IsAdministrator {
		//users := user.GetAllUsers()
		//moduleGroups := modulegroup.GetAllModuleGroups()

	} else {

	}

}
