package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetLatestSensorData(moduleGroupID int) (models.SensorData, error){
	db := database.GetDB()


}
