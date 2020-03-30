package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetAllModuleGroups() []models.ModuleGroup {
	db := database.GetDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM ModuleGroup;")
	if err != nil {
		panic(err)
	}

	var moduleGroups []models.ModuleGroup
	for rows.Next() {
		var id string
		var label string
		var plantID string
		var tds float32
		var ph float32
		var humidity float32
		var lightInterval []float32

		err := rows.Scan(&id, &label, &plantID, &tds, &ph, &humidity, &lightInterval)
		if err != nil {
			panic(err)
		}

		mg := models.ModuleGroup{
			ModuleGroupID:    id,
			ModuleGroupLabel: label,
			PlantID:          plantID,
			TDS:              tds,
			PH:               ph,
			Humidity:         humidity,
			LightInterval:    lightInterval,
		}

		moduleGroups = append(moduleGroups, mg)
	}

	return moduleGroups
}

func GetModuleGroup(moduleGroupID string) *models.ModuleGroup {
	db := database.GetDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM ModuleGroup WHERE ModuleGroupID = ?;", moduleGroupID)
	if err != nil {
		panic(err)
	}

	var moduleGroup *models.ModuleGroup
	for rows.Next() {
		var id string
		var label string
		var plantID string
		var tds float32
		var ph float32
		var humidity float32
		var lightsOn float32
		var lightsOff float32

		err := rows.Scan(&id, &label, &plantID, &tds, &ph, &humidity, &lightsOn, &lightsOff)
		if err != nil {
			panic(err)
		}

		moduleGroup = &models.ModuleGroup{
			ModuleGroupID:    id,
			ModuleGroupLabel: label,
			PlantID:          plantID,
			TDS:              tds,
			PH:               ph,
			Humidity:         humidity,
			LightsOn:         lightsOn,
			LightsOff:        lightsOff
		}
	}
	return moduleGroup
}

func NewModuleGroup(label string, plantID string, lightsOn float32, lightsOff float32) {
	db := database.GetDB()

	defer db.Close()

	plant := 

	sqlStatement := fmt.Sprintf("INSERT INTO ModuleGroup (ModuleGrouplabel, PlantiD, " +
		"TDS, PH, Humidity, LightsOn, LightsOff) VALUES (%s, %s, %g, %g, %g, %g)",
		)
	_, err := db.Exec(sqlStatement)

}
