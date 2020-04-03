package user_management

import (
	"fmt"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/user_management", getAllUsers)
	router.HandlerFunc("POST", "/user_management/register", user.CreateUser)

	return router
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
