package modulegroup_management

import (

	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", GetAllModuleGroup).
		Methods("GET").
		Schemes("http")

	return router
}