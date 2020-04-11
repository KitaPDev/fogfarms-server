package dashboard

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/dashboard", populateDashboard).
		Methods("GET").
		Schemes("http")
	router.HandleFunc("/dashboard/toggle", populateDashboard).
		Methods("GET").
		Schemes("http")

	return router
}

func populateDashboard(w http.ResponseWriter, r *http.Request) {
	PopulateDashboard(w, r)
}

func toggleDevice(w http.ResponseWriter, r *http.Request) {
	ToggleDevice(w, r)
}