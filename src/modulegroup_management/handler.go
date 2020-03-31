package modulegroup_management

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", getAllModuleGroup).
		Methods("GET").
		Schemes("http")

	return router
}

func getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}