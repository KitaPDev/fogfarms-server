package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/user"
)

const (
	DbHost = "localhost"
	DbPort = 5432
	DbUser = "fogfarms"
	DbPass = "fogfarms"
	DbName = "fogfarms-01"
)

func AssignUserToModuleGroup(username string, moduleGroupID string, level models.Level) {
	db := database.GetDB()
	u := user.GetUser(username)

	defer db.Close()
	sqlStatement := fmt.Sprintf("INSERT INTO Permission (Level, UserID, ModuleGroupID)" +
		"VALUES (%s, %s, %s)", string(level), u.UserID, moduleGroupID)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}