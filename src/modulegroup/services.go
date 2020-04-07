package modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup/repository"
)

func GetAllModuleGroups() ([]models.ModuleGroup, error) {
	moduleGroups, err := repository.GetAllModuleGroups()
	return moduleGroups, err
}

func GetModuleGroupsByID(moduleGroupIDs []int) ([]models.ModuleGroup, error) {
	moduleGroups, err := repository.GetModuleGroupsByID(moduleGroupIDs)
	return moduleGroups, err
}