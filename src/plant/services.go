package plant

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant/repository"
)

func GetAllPlants() ([]models.Plant, error) {
	return repository.GetAllPlants()
}

func NewPlant(plant models.Plant) error {
	return repository.NewPlant(plant.Name, plant.TDS, plant.PH, plant.Lux, plant.LightsOnHour,
		plant.LightsOffHour)
}

func DeletePlant(plantID int) error {
	return repository.DeletePlant(plantID)
}