package modulegroup

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllModuleGroups() ([]models.ModuleGroup, error)
	GetModuleGroupByID(moduleGroupID int) (*models.ModuleGroup, error)
	GetModuleGroupsByID(moduleGroupIDs []int) error
	NewModuleGroup(moduleGroupLabel string, plantID int, lightsOn float32, lightsOff float32) error
	ToggleAuto(moduleGroupID int) error
	SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
		lightsOn float32, lightsOff float32) error
	ResetTimer(moduleGroupID int) error
}