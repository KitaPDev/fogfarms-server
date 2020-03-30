package plant

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetPlant(plantID string) *models.Plant
	GetAllPlants() []models.Plant
	NewPlant(name string, tds float32, ph float32, lux float32)
}