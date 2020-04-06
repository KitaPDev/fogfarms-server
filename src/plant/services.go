package plant

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant/repository"
)

func GetPlant(plantID int) *models.Plant {
	return repository.GetPlant(plantID)
}