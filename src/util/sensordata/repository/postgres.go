package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
	"github.com/lib/pq"
)

func GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error) {
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

	db := database.GetDB()

	sqlDropTempTable := `DROP TABLE IF EXISTS temp_SensorData`
	_, err = db.Exec(sqlDropTempTable)
	if err != nil {
		return nil, err
	}

	sqlCreateFunction :=
		`CREATE OR REPLACE FUNCTION fn_getLatestSensorData(moduleIDs INTEGER ARRAY)
			RETURNS SETOF SensorData AS
		$func$
			DECLARE
				i INT;
		
			BEGIN
				CREATE TEMPORARY TABLE temp_SensorData (
				   ModuleID INT,
				   Timestamp TIMESTAMP,
				   TDS FLOAT,
				   PH FLOAT,
				   SolutionTemperature FLOAT,
				   ArrGrowUnitLux FLOAT ARRAY,
				   ArrGrowUnitHumidity FLOAT ARRAY,
				   ArrGrowUnitTemperature FLOAT ARRAY
				);
		
				FOREACH i IN ARRAY moduleIDs
					LOOP
						INSERT INTO temp_SensorData (ModuleID, Timestamp, TDS, PH, SolutionTemperature, ArrGrowUnitLux, ArrGrowUnitHumidity, ArrGrowUnitTemperature)
						SELECT * FROM SensorData WHERE ModuleID = i ORDER BY Timestamp DESC LIMIT 1;
					END LOOP;
		
				RETURN QUERY SELECT * FROM temp_SensorData;
			END
		$func$ LANGUAGE plpgsql;`

	_, err = db.Exec(sqlCreateFunction)
	if err != nil {
		return nil, err
	}

	sqlStatement := `SELECT * FROM fn_getLatestSensorData($1);`

	rows, err := db.Query(sqlStatement, pq.Array(moduleIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

	return sensorData, nil
}