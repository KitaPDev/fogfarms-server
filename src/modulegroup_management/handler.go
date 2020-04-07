package modulegroup_management

import (
	"fmt"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", GetAllModuleGroup).
		Methods("GET").
		Schemes("http")
	router.HandleFunc("/modulegroup_management/test", test).
		Methods("GET").
		Schemes("http")

	return router
}
func test(w http.ResponseWriter, r *http.Request) {
	v := jwt.AuthenticateUserToken(w, r)
	fmt.Printf(("%+v"), v)
}
