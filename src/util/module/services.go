package module

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/module/repository"
)

func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	return repository.GetModulesByModuleGroupIDs(moduleGroupIDs)
}