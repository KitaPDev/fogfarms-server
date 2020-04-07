package plant

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant/repository"
)

func GetAllPlants() ([]models.Plant, error) {
	plant, err := repository.GetAllPlants()
	return plant, err
}

func NewPlant(plant models.Plant) error {
	return repository.NewPlant(plant.Name, plant.TDS, plant.PH, plant.Lux, plant.LightsOnHour,
		plant.LightsOffHour)
}

func DeletePlant(plantID int) error {
	return repository.DeletePlant(plantID)
}

func GetPlantByID(plantID int) (*models.Plant, error) {
	plant, err := repository.GetPlantByID(plantID)
	return plant, err
}
