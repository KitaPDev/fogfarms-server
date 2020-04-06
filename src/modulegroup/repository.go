package modulegroup

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllModuleGroups() []models.ModuleGroup
	GetModuleGroupByID(moduleGroupID int) *models.ModuleGroup
	GetModuleGroupsByID(moduleGroupIDs []int)
	NewModuleGroup(moduleGroupLabel string, plantID int, lightsOn float32, lightsOff float32)
	SetManualOperation(moduleGroupID int, toManual bool)
	SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
		lightsOn float32, lightsOff float32)
}