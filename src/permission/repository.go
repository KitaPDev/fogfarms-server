package permission

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllPermissions () []models.Permission
	AssignUserToModuleGroup(username string, moduleGroupID string, permissionLevel int)
	GetSupervisorModuleGroups(userID int) []models.ModuleGroup
}