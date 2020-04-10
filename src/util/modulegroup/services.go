package modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup/repository"
)

func GetAllModuleGroups() ([]models.ModuleGroup, error) {
	moduleGroups, err := repository.GetAllModuleGroups()
	return moduleGroups, err
}

func GetModuleGroupsByIDs(moduleGroupIDs []int) ([]models.ModuleGroup, error) {
	moduleGroups, err := repository.GetModuleGroupsByIDs(moduleGroupIDs)
	return moduleGroups, err
}