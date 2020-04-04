package permission

import (
	"github.com/KitaPDev/fogfarms-server/models"
)

func AssignUserToModuleGroup(username string, moduleGroupID string, permissionLevel int) {
	//repository.AssignUserToModuleGroup(username, moduleGroupID, permissionLevel)
}

func GetSupervisorModuleGroups(user *models.User) []models.ModuleGroup {
	return nil
}