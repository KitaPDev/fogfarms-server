package modulegroup_management

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/modulegroup_management", GetAllModuleGroup)
	router.HandlerFunc("GET", "/modulegroup_management/js", GetDemoJson)
	router.HandlerFunc("GET", "/modulegroup_management/db", GetTestName)
	router.HandlerFunc("POST", "/modulegroup_management/db", PostTestName)
	return router
}
