package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
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
	user := user.GetUser(username)

}