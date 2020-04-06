package plant_management

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant/repository"
	"log"
	"net/http"
)

func GetAllPlants(w http.ResponseWriter) {
	plants, err := repository.GetAllPlants()
	if err != nil {
		msg := "Failed to GetAllPlants"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	plantsJson, err := json.Marshal(plants)
	if err != nil {
		msg := "Failed to Marshal plants to JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Fatal(w.Write(plantsJson))
	return
}

func NewPlant(w http.ResponseWriter, r *http.Request) {
	input := models.Plant{}


}