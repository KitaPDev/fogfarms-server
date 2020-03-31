package plant_management

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/plant_management", getAllPlants).
		Methods("GET").
		Schemes("http")

	return router
}

func getAllPlants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}