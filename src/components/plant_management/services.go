package plant_management

import (
	"encoding/json"
	"fmt"
	"github.com/KitaPDev/fogfarms-server/src/util/auth/jwt"
	"log"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/plant"
	"github.com/golang/gddo/httputil/header"
)

func GetAllPlants(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	plants, err := plant.GetAllPlants()
	fmt.Printf("Hi")
	if err != nil {
		msg := "Error: Failed to Get All Plants"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
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
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(plantsJson)
}

func CreatePlant(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

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
		msg := "Error: Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = plant.CreatePlant(input)
	if err != nil {
		msg := "Error: Failed to New Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: New Plant; Successful"))
}

func DeletePlant(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

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
		msg := "Error: Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = plant.DeletePlant(input.PlantID)
	if err != nil {
		msg := "Error: Failed to Delete Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Delete Plant; Successful"))
}
