package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetPlant(plantID string) *models.Plant {
	db := database.GetDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM Plant WHERE PlantID = ?", plantID)
	if err != nil {
		panic(err)
	}

	var plant *models.Plant
	for rows.Next() {
		var id string
		var name string
		var tds float32
		var ph float32
		var lux float32

		err := rows.Scan(&id, &name, &tds, &ph, &lux)
		if err != nil {
			panic(err)
		}

		plant = &models.Plant{
			PlantID: id,
			Name:    name,
			TDS:     tds,
			PH:      ph,
			Lux:     lux,
		}
	}
	return plant
}

func GetAllPlants() []models.Plant {
	db := database.GetDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM Plant")
	if err != nil {
		panic(err)
	}

	var plants []models.Plant
	for rows.Next() {
		var id string
		var name string
		var tds float32
		var ph float32
		var lux float32

		err := rows.Scan(&id, &name, &tds, &ph, &lux)
		if err != nil {
			panic(err)
		}

		plant := models.Plant{
			PlantID: id,
			Name:    name,
			TDS:     tds,
			PH:      ph,
			Lux:     lux,
		}

		plants = append(plants, plant)
	}
	return plants
}

func NewPlant(name string, tds float32, ph float32, lux float32) {
	db := database.GetDB()

	defer db.Close()
	sqlStatement := fmt.Sprintf("INSERT INTO Plant (Name, TDS, PH, Lux)" +
		"VALUES (%s, %g, %g, %g)", name, tds, ph, lux)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}