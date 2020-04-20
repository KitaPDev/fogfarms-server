package repository

import (
	"log"
	"math/rand"
	"time"

	"github.com/KitaPDev/fogfarms-server/models/outputs"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/lib/pq"
)

func CreateModule(moduleLabel string) error {
	db := database.GetDB()

	sqlStatement :=
		`INSERT INTO Module (ModuleLabel, Token, ArrFogger, ArrLED, ArrMixer, ArrSolenoidValve)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Query(sqlStatement, moduleLabel, GenerateToken(), pq.BoolArray{}, pq.BoolArray{},
				pq.BoolArray{}, pq.BoolArray{})
	if err != nil {
		return err
	}

	return nil
}

func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ModuleID, ModuleGroupID, ModuleLabel FROM Module WHERE ModuleGroupID = ANY($1)`

	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []models.Module
	for rows.Next() {
		module := models.Module{}

		err := rows.Scan(&module.ModuleID, &module.ModuleGroupID, &module.ModuleLabel)
		if err != nil {
			return nil, err
		}

		modules = append(modules, module)
	}
	return modules, nil
}

func GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs []int) ([]outputs.ModuleOutput, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT module.moduleID, module.moduleGroupID, modulelabel,modulegrouplabel FROM Module,Modulegroup WHERE module.ModuleGroupID = ANY($1) AND module.modulegroupID=modulegroup.modulegroupID ;`

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

func AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE Module SET ModuleGroupID = $1 WHERE ModuleID = ANY($2)`

	_, err := db.Query(sqlStatement, moduleGroupID, pq.Array(moduleIDs))

	return err
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}