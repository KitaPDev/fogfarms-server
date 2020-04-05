package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/plant"
	"github.com/labstack/gommon/log"
)

func GetAllModuleGroups() []models.ModuleGroup {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM ModuleGroup;")
	if err != nil {
		panic(err)
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
			&moduleGroup.LightsOnTime,
			&moduleGroup.LightsOffTime,
		)
		if err != nil {
			panic(err)
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	return moduleGroups
}

func GetModuleGroup(moduleGroupID string) *models.ModuleGroup {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM ModuleGroup WHERE ModuleGroupID = ?;", moduleGroupID)
	if err != nil {
		panic(err)
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
			&moduleGroup.LightsOnTime,
			&moduleGroup.LightsOffTime,
		)
		if err != nil {
			panic(err)
		}
	}

	return moduleGroup
}

func NewModuleGroup(label string, plantID string, humidity float32, lightsOn float32,
	lightsOff float32) {
	db := database.GetDB()

	plant := plant.GetPlant(plantID)

	sqlStatement := fmt.Sprintf("INSERT INTO ModuleGroup (ModuleGrouplabel, PlantID, "+
		"TDS, PH, Humidity, LightsOn, LightsOff) VALUES (%s, %s, %g, %g, %g, %g, %g)",
		label, plantID, plant.TDS, plant.PH, humidity, lightsOn, lightsOff)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func SetManualOperation(moduleGroupID string, toManual bool) {
	db := database.GetDB()

	sqlStatement := fmt.Sprintf("UPDATE ModuleGroup SET OnAuto = %t " +
		"WHERE ModuleGroupID = %s", toManual, moduleGroupID)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func SetEnvironmentParameters(moduleGroupID string, humidity float32, ph float32, tds float32,
	lightsOn float32, lightsOff float32) {
	db := database.GetDB()

	sqlStatement := fmt.Sprintf("UPDATE ModuleGroup" +
		"SET Humidity = %g, PH = %g, TDS = %g, LightsOn = %g, LightsOff = %g" +
		"WHERE ModuleGroupID = %s", humidity, ph, tds, lightsOn, lightsOff, moduleGroupID)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}