package plant_management

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/plant_management", getAllPlants).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/plant_management/new_plant", newPlant).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/plant_management/delete_plant", deletePlant).
		Methods("GET").
		Schemes("http")

	return router
}

// todo ddfsd test
func getAllPlants(w http.ResponseWriter) {
	GetAllPlants(w)
}

// todo ddfsd test
func newPlant(w http.ResponseWriter, r *http.Request) {
	NewPlant(w, r)
}

// todo ddfsd test
func deletePlant(w http.ResponseWriter, r *http.Request) {
	DeletePlant(w, r)
}