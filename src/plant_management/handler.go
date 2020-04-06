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

func getAllPlants(w http.ResponseWriter, r *http.Request) {
	GetAllPlants()
}

func newPlant(w http.ResponseWriter, r *http.Request) {

}

func deletePlant(w http.ResponseWriter, r *http.Request) {

}