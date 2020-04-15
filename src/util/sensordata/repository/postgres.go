package repository

import (
	"log"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
	"github.com/lib/pq"
)

func GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error) {
	var moduleGroupIDs []int
	moduleGroupIDs = append(moduleGroupIDs, moduleGroupID)
	log.Println("Modulegroupids in getlatest sensor", moduleGroupIDs)
	modules, err := module.GetModulesByModuleGroupIDs(moduleGroupIDs)
	if err != nil {
		return nil, err
	}
	var moduleIDs []int
	for _, m := range modules {
		moduleIDs = append(moduleIDs, m.ModuleID)
	}

	log.Println("Moduleids in getlatest sensor", moduleIDs)

	sqlStatement := `SELECT * FROM SensorData WHERE ModuleID = ANY($1)`

	db := database.GetDB()
	rows, err := db.Query(sqlStatement, pq.Array(moduleIDs))
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
			pq.Array(&sd.GrowUnitLux),
			pq.Array(&sd.GrowUnitHumidity),
			pq.Array(&sd.GrowUnitTemperature),
		)
		if err != nil {
			return nil, err
		}
		sensorData = append(sensorData, sd)

	}
	log.Println("sensordata in getlatest sensor", sensorData)

	return sensorData, nil
}
