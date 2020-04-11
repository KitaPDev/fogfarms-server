package repository

import "github.com/KitaPDev/fogfarms-server/models"
import "github.com/KitaPDev/fogfarms-server/src/database"

func GetModuleGroupDevices(moduleGroupID int) ([]models.Device, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT * FROM Device WHERE ModuleID 
		IN (SELECT ModuleID FROM ModuleGroup WHERE ModuleGroupID = $1)`

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var d models.Device

		err = rows.Scan(&d)
		if err != nil {
			return nil, err
		}

		devices = append(devices, d)
	}

	return devices, nil
}