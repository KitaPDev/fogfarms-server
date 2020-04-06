package permission

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllPermissions () []models.Permission
	AssignUserToModuleGroup(userID int, moduleGroupID int, permissionLevel int)
	GetSupervisorModuleGroups(userID int) []models.ModuleGroup
}