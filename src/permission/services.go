package permission

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/permission/repository"
)

func AssignUserToModuleGroup(username string, moduleGroupID string, level string) {
	var l models.Level
	if level == "Supervisor" || level == "Control" || level == "Monitor" {
		l = models.Level(level)
	}

	repository.AssignUserToModuleGroup(username, moduleGroupID, l)
}