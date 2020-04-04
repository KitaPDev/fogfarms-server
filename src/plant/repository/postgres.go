package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"log"
)

func GetPlant(plantID string) *models.Plant {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Plant WHERE PlantID = ?", plantID)
	if err != nil {
		panic(err)
	}
	defer log.Fatal(rows.Close())

	plant := &models.Plant{}
	for rows.Next() {

		err := rows.Scan(
			&plant.PlantID,
			&plant.Name,
			&plant.TDS,
			&plant.PH,
			&plant.Lux,
		)
		if err != nil {
			panic(err)
		}
	}
	return plant
}

func GetAllPlants() []models.Plant {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Plant")
	if err != nil {
		panic(err)
	}
	defer log.Fatal(rows.Close())

	var plants []models.Plant
	for rows.Next() {
		plant := &models.Plant{}

		err := rows.Scan(
			&plant.PlantID,
			&plant.Name,
			&plant.TDS,
			&plant.PH,
			&plant.Lux,
		)
		if err != nil {
			panic(err)
		}

		plants = append(plants, plant)
	}
	return plants
}

func NewPlant(name string, tds float32, ph float32, lux float32) {
	db := database.GetDB()

	sqlStatement := fmt.Sprintf("INSERT INTO Plant (Name, TDS, PH, Lux)" +
		"VALUES (%s, %g, %g, %g)", name, tds, ph, lux)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}