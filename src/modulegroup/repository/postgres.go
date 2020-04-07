package repository

import (
	"time"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/plant"
	"github.com/labstack/gommon/log"
)

func GetAllModuleGroups() ([]models.ModuleGroup, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM ModuleGroup;")
	if err != nil {
		return nil, err
	}
	defer log.Fatal(rows.Close())

	var moduleGroups []models.ModuleGroup
	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
		)
		if err != nil {
			return nil, err
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	return moduleGroups, nil
}

func GetModuleGroupByID(moduleGroupID int) (*models.ModuleGroup, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM ModuleGroup WHERE ModuleGroupID = ?;", moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer log.Fatal(rows.Close())

	moduleGroup := &models.ModuleGroup{}
	for rows.Next() {

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
		)
		if err != nil {
			return nil, err
		}
	}

	return moduleGroup, nil
}

func GetModuleGroupsByID(moduleGroupIDs []int) ([]models.ModuleGroup, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM ModuleGroup WHERE ModuleGroupID IN (?);", moduleGroupIDs)
	if err != nil {
		return nil, err
	}
	defer log.Fatal(rows.Close())

	var moduleGroups []models.ModuleGroup
	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
		)
		if err != nil {
			return nil, err
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	return moduleGroups, err
}

func NewModuleGroup(label string, plantID int, locationID int, humidity float32, lightsOn float32,
	lightsOff float32) {
	db := database.GetDB()

	p, errTemp := plant.GetPlantByID(plantID)
	if errTemp != nil {
		panic(errTemp)
	}
	sqlStatement := `INSERT INTO ModuleGroup (ModuleGroupLabel, PlantID, LocationID,
                         Param_TDS, Param_Ph, Param_Humidity, LightsOnHour, LightsOffHour)
                         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Query(sqlStatement, label, plantID, locationID, p.TDS, p.PH, humidity, lightsOn, lightsOff)
	if err != nil {
		panic(err)
	}
}

func SetManualOperation(moduleGroupID int, toManual bool) {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET OnAuto = $1
		WHERE ModuleGroupID = $2`
	_, err := db.Query(sqlStatement, toManual, moduleGroupID)
	if err != nil {
		panic(err)
	}
}

func SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
	lightsOn time.Time, lightsOff time.Time) {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup	
						SET param_humidity = $1, param_ph = $2, param_tds = $3, lightsofftime = $4, 
						    lightsontime = $5
						WHERE ModuleGroupID = $6`
	_, err := db.Query(sqlStatement, humidity, ph, tds, lightsOff, lightsOn, moduleGroupID)
	if err != nil {
		panic(err)
	}
}
