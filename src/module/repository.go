package module

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error)
}

