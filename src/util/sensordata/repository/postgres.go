package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
)

func GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error){
	var moduleGroupIDs []int
	moduleGroupIDs = append(moduleGroupIDs, moduleGroupID)

	modules, err := module.GetModulesByModuleGroupIDs(moduleGroupIDs)
	if err != nil {
		return nil, err
	}

	var moduleIDs []int
	for _, m := range modules {
		moduleIDs = append(moduleIDs, m.ModuleID)
	}

	sqlStatement := `SELECT * FROM SensorData WHERE ModuleID = ANY ($1)`

	db := database.GetDB()
	rows, err := db.Query(sqlStatement, moduleIDs)
	if err != nil {
		return nil, err
	}

	var sensorData []models.SensorData
	for rows.Next() {
		var sd models.SensorData

		err = rows.Scan(
			&sd.ModuleID,
			&sd.TimeStamp,
			&sd.TDS,
			&sd.PH,
			&sd.SolutionTemperature,
			&sd.GrowUnitLux,
			&sd.GrowUnitHumidity,
			&sd.GrowUnitTemperature,
		)
		if err != nil {
			return nil, err
		}
	}

	return sensorData, nil
}
