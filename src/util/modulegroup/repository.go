package modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"time"
)

type Repository interface {
	GetAllModuleGroups() ([]models.ModuleGroup, error)
	GetModuleGroupByID(moduleGroupID int) (*models.ModuleGroup, error)
	GetModuleGroupsByID(moduleGroupIDs []int) error
	CreateModuleGroup(label string, plantID int, locationID int, humidity float32, lightsOn float32,
		lightsOff float32, onAuto bool, timerLastReset time.Time) error
	NewModuleGroup(moduleGroupLabel string, plantID int, lightsOn float32, lightsOff float32) error
	ToggleAuto(moduleGroupID int) error
	SetEnvironmentParameters(moduleGroupID int, humidity float32, ph float32, tds float32,
		lightsOn float32, lightsOff float32) error
	AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error
	ResetTimer(moduleGroupID int) error
}