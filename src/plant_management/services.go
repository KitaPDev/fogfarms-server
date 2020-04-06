package plant_management

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func GetAllPlants(w http.ResponseWriter) {
	plants, err := plant.GetAllPlants()
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
}

func NewPlant(w http.ResponseWriter, r *http.Request) {
	input := models.Plant{}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		msg := "Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	err = plant.NewPlant(input)
	if err != nil {
		msg := "Failed to Create New Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeletePlant(w http.ResponseWriter, r *http.Request) {
	input := models.Plant{}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		msg := "Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	err = plant.DeletePlant(input.PlantID)
	if err != nil {
		msg := "Failed to Delete Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}