package modulegroup_management

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", populateModuleGroupManagementPage).
		Methods("GET").
		Schemes("http")

	return router
}

func populateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {

}