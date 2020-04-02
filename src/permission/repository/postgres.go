package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"log"
)

func AssignUserToModuleGroup(username string, moduleGroupID string, level models.Level) {
	db := database.GetDB()
	u := user.GetUser(username)

	sqlStatement := fmt.Sprintf("INSERT INTO Permission (Level, UserID, ModuleGroupID)" +
		"VALUES (%s, %s, %s)", string(level), u.UserID, moduleGroupID)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func GetSupervisorModuleGroup(userID string) []models.ModuleGroup {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Permission")
	defer log.Fatal(rows.Close())

} 