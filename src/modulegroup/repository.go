package modulegroup

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllModuleGroups() []models.ModuleGroup
	GetModuleGroup(moduleGroupID string) *models.ModuleGroup
	NewModuleGroup(moduleGroupLabel string, plantID string, lightsOn float32, lightsOff float32)
	SetManualOperation(moduleGroupID string, toManual bool)
	SetEnvironmentParameters(moduleGroupID string, humidity float32, ph float32, tds float32,
		lightsOn float32, lightsOff float32)
}