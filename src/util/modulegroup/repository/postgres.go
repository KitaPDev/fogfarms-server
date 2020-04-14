package repository

import (
	"log"
	"time"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/util/plant"
	"github.com/lib/pq"
)

func GetAllModuleGroups() ([]models.ModuleGroup, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, Param_TDs, Param_PH, 
		Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset FROM ModuleGroup;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			&moduleGroup.TimerLastReset,
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

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, Param_TDs, Param_PH, 
		Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset
		FROM ModuleGroup WHERE ModuleGroupID = $1;`

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			return nil, err
		}
	}

	return moduleGroup, nil
}

func GetModuleGroupsByIDs(moduleGroupIDs []int) ([]models.ModuleGroup, error) {
	var moduleGroups []models.ModuleGroup
	var err error

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, Param_TDs, Param_PH, 
		Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset
		FROM ModuleGroup WHERE ModuleGroupID = ANY($1);`

	db := database.GetDB()

	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.LocationID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.OnAuto,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			return nil, err
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	log.Println("Variable moduleGroups in GetModuleGroups by ID", moduleGroups)

	return moduleGroups, err
}

func CreateModuleGroup(label string, plantID int, locationID int, humidity float64, lightsOn float64,
	lightsOff float64, onAuto bool, timerLastReset time.Time) error {
	db := database.GetDB()

	p, err := plant.GetPlantByID(plantID)
	if err != nil {
		return err
	}
	sqlStatement :=
		`INSERT INTO ModuleGroup (ModuleGroupLabel, PlantID, LocationID, onAuto,
		 Param_TDS, Param_Ph, Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, Now())`
	_, err = db.Query(sqlStatement, label, plantID, locationID, onAuto, p.TDS, p.PH, humidity, lightsOn, lightsOff)
	if err != nil {
		return err
	}

	return nil
}

func ToggleAuto(moduleGroupID int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET OnAuto = NOT OnAuto WHERE ModuleGroupID = $1`
	_, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
	lightsOnHour float32, lightsOffHour float32) error {
	db := database.GetDB()

	sqlStatement :=
		`UPDATE ModuleGroup	
			SET Param_Humidity = $1, Param_PH = $2, Param_TDS = $3, LightsOnHour = $4, 
				LightsOffHour = $5
			WHERE ModuleGroupID = $6`
	_, err := db.Query(sqlStatement, humidity, ph, tds, lightsOffHour, lightsOnHour, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE Module SET ModuleGroupID = $1 WHERE ModuleID = ANY($2)`

	_, err := db.Query(sqlStatement, moduleGroupID, moduleIDs)

	return err
}

func ResetTimer(moduleGroupID int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET TimerLastReset = NOW() WHERE ModuleGroupID = $1;`

	_, err := db.Query(sqlStatement, moduleGroupID)

	return err
}
