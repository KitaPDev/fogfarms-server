package modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
)

type Repository interface {
	GetAllModuleGroups() ([]models.ModuleGroup, error)
	GetModuleGroupByID(moduleGroupID int) (*models.ModuleGroup, error)
	GetModuleGroupsByID(moduleGroupIDs []int) error
	CreateModuleGroup(label string, plantID int, locationID int, humidity float64, lightsOn float64,
		lightsOff float64, onAuto bool) error
	NewModuleGroup(moduleGroupLabel string, plantID int, lightsOn float64, lightsOff float64) error
	ToggleAuto(moduleGroupID int) error
	SetEnvironmentParameters(moduleGroupID int, humidity float64, ph float64, tds float64,
		lightsOn float64, lightsOff float64) error
	AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error
	ResetTimer(moduleGroupID int) error
}