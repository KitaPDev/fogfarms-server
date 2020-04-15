package repository

import (
	"log"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetModuleGroupDevices(moduleGroupID int) ([]models.Device, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT Distinct DeviceID,DeviceTypeID,IsOn,GrowUnitID,NutrientUnitID,PHDownUnitID,PHUpUnitID FROM Device 
		WHERE Device.Growunitid IN (SELECT growunit.growunitid FROM growunit WHERE growunit.ModuleID IN (SELECT module.moduleid from module WHERE module.modulegroupid = $1))
		OR Device.nutrientunitid IN (SELECT nutrientunit.nutrientunitid FROM nutrientunit WHERE nutrientunit.ModuleID IN (SELECT module.moduleid from module WHERE module.modulegroupid = $1))`

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var d models.Device

		err = rows.Scan(&d.DeviceID, &d.DeviceTypeID, &d.IsOn, &d.GrowUnitID, &d.NutrientUnitID, &d.PHDownUnitID, &d.PHUpUnitID)
		if err != nil {
			return nil, err
		}
		log.Print("device is fine")
		devices = append(devices, d)
	}

	return devices, nil
}

func ToggleDevice(deviceID int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE Device SET IsOn = NOT IsOn WHERE DeviceID = $1`

	_, err := db.Query(sqlStatement, deviceID)
	return err
}
