package modulegroup

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllModuleGroups() ([]models.ModuleGroup, error)
	GetModuleGroupByID(moduleGroupID int) (*models.ModuleGroup, error)
	GetModuleGroupsByID(moduleGroupIDs []int) error
	NewModuleGroup(moduleGroupLabel string, plantID int, lightsOn float32, lightsOff float32) error
	SetManualOperation(moduleGroupID int, toManual bool) error
	SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
		lightsOn float32, lightsOff float32) error
}