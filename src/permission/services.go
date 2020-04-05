package permission

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/permission/repository"
)

func AssignUserToModuleGroup(username string, moduleGroupID int, permissionLevel int) {
	repository.AssignUserToModuleGroup(username, moduleGroupID, permissionLevel)
}

func GetSupervisorModuleGroups(user *models.User) []models.ModuleGroup {
	return nil
}