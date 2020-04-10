package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT * FROM Module
		WHERE ModuleGroupID = ANY($1)`

	rows, err := db.Query(sqlStatement, moduleGroupIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []models.Module
	for rows.Next() {
		module := models.Module{}

		err := rows.Scan(&module)
		if err != nil {
			return nil, err
		}

		modules = append(modules, module)
	}

	return modules, nil
}