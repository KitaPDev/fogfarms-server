package repository

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"log"
)

func AssignUserToModuleGroup(username string, moduleGroupID string, level int) {
	db := database.GetDB()
	u := user.GetUser(username)

	sqlStatement := fmt.Sprintf("INSERT INTO Permission (PermissionLevel, UserID, ModuleGroupID)" +
		"VALUES (%d, %s, %s)", level, u.UserID, moduleGroupID)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func GetSupervisorModuleGroups(userID string) []models.ModuleGroup {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Permission WHERE UserID = ?", userID)
	if err != nil {
		panic(err)
	}
	defer log.Fatal(rows.Close())

	var moduleGroups []models.ModuleGroup
	for rows.Next() {

		//err := rows.Scan()

	}

	return moduleGroups
} 