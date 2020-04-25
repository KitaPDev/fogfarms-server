package sensordata

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/models/outputs"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata/repository"
	"time"
)

func GetLatestSensorData(moduleGroupID int) (map[string]*outputs.DashboardOutput, error) {
	return repository.GetLatestSensorData(moduleGroupID)
}

func GetSensorDataHistory(moduleGroupID int, timeBegin time.Time, timeEnd time.Time) (map[string][]models.SensorData, error) {
	return repository.GetSensorDataHistory(moduleGroupID, timeBegin, timeEnd)
}

func RecordSensorData(moduleID int, tds []float64, ph []float64, solutionTemperature []float64,
	lux []float64, humidity []float64, temperature []float64) error {

	return repository.RecordSensorData(moduleID, tds, ph, solutionTemperature, lux, humidity, temperature)
}