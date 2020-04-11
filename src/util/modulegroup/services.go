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

func ToggleAuto(moduleGroupID int) error {
	return repository.ToggleAuto(moduleGroupID)
}

func SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
	lightsOnHour float32, lightsOffHour float32) error {

	return repository.SetEnvironmentParameters(moduleGroupID, humidity, ph, tds, lightsOnHour,
		lightsOffHour)
}

func ResetTimer(moduleGroupID int) error {
	return repository.ResetTimer(moduleGroupID)
}