package modulegroup_management

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", populateModuleGroupManagementPage).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/create", createModuleGroup).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/assign", assignModuleToModuleGroup).
		Methods("POST").
		Schemes("http")

	return router
}

func populateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	PopulateModuleGroupManagementPage(w, r)
}

func createModuleGroup(w http.ResponseWriter, r *http.Request) {
	CreateModuleGroup(w, r)
}

func assignModuleToModuleGroup(w http.ResponseWriter, r *http.Request) {
	AssignModuleToModuleGroup(w, r)
}