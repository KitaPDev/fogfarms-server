package plant

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant/repository"
)

func GetPlant(plantID string) *models.Plant {
	return repository.GetPlant(plantID)
}