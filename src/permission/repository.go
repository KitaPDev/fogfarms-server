package permission

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	AssignUserToModuleGroup(username string, moduleGroupID string, level models.Level)
}