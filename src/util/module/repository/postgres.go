package repository

import (
	"log"

	"github.com/KitaPDev/fogfarms-server/models/outputs"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/lib/pq"
)

func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT moduleid, moduleGroupID, modulelabel FROM Module WHERE ModuleGroupID = ANY($1) ;`
	log.Println(sqlStatement)
	log.Println(moduleGroupIDs)
	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []models.Module
	for rows.Next() {
		module := models.Module{}
		log.Println(module)
		err := rows.Scan(&module.ModuleID, &module.ModuleGroupID, &module.ModuleLabel)
		if err != nil {
			return nil, err
		}
		log.Println(module)
		modules = append(modules, module)
	}
	return modules, nil
}

func GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs []int) ([]outputs.ModuleOutput, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT module.moduleID, module.moduleGroupID, modulelabel,modulegrouplabel FROM Module,Modulegroup WHERE module.ModuleGroupID = ANY($1) AND module.modulegroupID=modulegroup.modulegroupID ;`
	log.Println(sqlStatement)
	log.Println(moduleGroupIDs)
	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []outputs.ModuleOutput
	for rows.Next() {
		module := outputs.ModuleOutput{}
		log.Println(module)
		err := rows.Scan(&module.ModuleID, &module.ModuleGroupID, &module.ModuleLabel, &module.ModuleGroupLabel)
		if err != nil {
			return nil, err
		}
		log.Println(module)
		modules = append(modules, module)
	}
	return modules, nil
}
