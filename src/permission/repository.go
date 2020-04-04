package permission

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	AssignUserToModuleGroup(username string, moduleGroupID string, permissionLevel int)
	GetSupervisorModuleGroups(userID string) []models.ModuleGroup
}