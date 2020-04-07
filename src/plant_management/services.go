package plant_management

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant"
	"github.com/golang/gddo/httputil/header"
)

func GetAllPlants(w http.ResponseWriter, r *http.Request) {
	plants, err := plant.GetAllPlants()
	fmt.Printf("Hi")
	if err != nil {
		msg := "Error: plant.GetAllPlants()"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v", plants)
	type Output struct {
		Data []models.Plant
	}
	out := Output{plants}
	plantsJson, err := json.Marshal(out)
	if err != nil {
		msg := "Error: json.Marshal(plants)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(plantsJson)
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
		msg := "Error: json.NewDecoder(r.Body).Decode(&input)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	fmt.Print("%+v", input)
	err = plant.NewPlant(input)
	if err != nil {
		msg := "Error: plant.NewPlant(input)"
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
		msg := "Error: json.NewDecoder(r.Body).Decode(&input)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	err = plant.DeletePlant(input.PlantID)
	if err != nil {
		msg := "Error: plant.DeletePlant(input.PlantID)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
