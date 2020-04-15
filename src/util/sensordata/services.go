package sensordata

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata/repository"
)

func GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error) {
	return repository.GetLatestSensorData(moduleGroupID)
}