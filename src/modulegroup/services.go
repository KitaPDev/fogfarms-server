package modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup/repository"
)

func GetAllModuleGroups() []models.ModuleGroup{
	return repository.GetAllModuleGroups()
}

func GetModuleGroupsByID(moduleGroupIDs []int) []models.ModuleGroup {

}