package module

import (
	"github.com/KitaPDev/fogfarms-server/models"
)

type Repository interface {
	CreateModule(moduleLabel string) error
	GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error)
	AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error
}

