package modulegroup_management

import (
	"fmt"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", populateModuleGroupManagementPage).
		Methods("GET").
		Schemes("http")
	router.HandleFunc("/modulegroup_management/test", test).
		Methods("GET").
		Schemes("http")

	return router
}

func populateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {

}
