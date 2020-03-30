package plant

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/plant"
)

func GetPlant(plantID string) *models.Plant {
	plant.GetPlant(plantID)
}