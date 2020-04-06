package plant

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant/repository"
)

func GetAllPlants() ([]models.Plant, error) {
	plants, err := repository.GetAllPlants()
	if err != nil {
		return nil, err
	}

	return plants, nil
}

func NewPlant(plant models.Plant) error {
	err := repository.NewPlant(plant.Name, plant.TDS, plant.PH, plant.Lux, plant.LightsOnHour,
		plant.LightsOffHour)
	if err != nil {
		return err
	}

	return nil
}

func DeletePlant(plantID int) error {
	err := repository.DeletePlant(plantID)
	if err != nil {
		return err
	}

	return nil
}