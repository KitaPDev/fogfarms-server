package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"log"
)

func GetPlant(plantID int) (*models.Plant, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Plant WHERE PlantID = ?", plantID)
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}
	return plant, nil
}

func GetAllPlants() ([]models.Plant, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Plant")
	if err != nil {
		return nil, err
	}
	defer log.Fatal(rows.Close())

	var plants []models.Plant
	for rows.Next() {
		plant := models.Plant{}

		err := rows.Scan(
			&plant.PlantID,
			&plant.Name,
			&plant.TDS,
			&plant.PH,
			&plant.Lux,
		)
		if err != nil {
			return nil, err
		}

		plants = append(plants, plant)
	}
	return plants, nil
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