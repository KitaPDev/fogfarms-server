package plant_management

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/plant_management", getAllPlants).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/plant_management/new_plant", newPlant).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/plant_management/delete_plant", deletePlant).
		Methods("POST").
		Schemes("http")

	return router
}

// ok
func getAllPlants(w http.ResponseWriter, r *http.Request) {
	GetAllPlants(w, r)
}

//  ok
func newPlant(w http.ResponseWriter, r *http.Request) {
	NewPlant(w, r)
}

// ok
func deletePlant(w http.ResponseWriter, r *http.Request) {
	DeletePlant(w, r)
}
