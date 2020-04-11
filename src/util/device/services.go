package device

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/device/repository"
)

func GetModuleGroupDevices(moduleGroupID int) ([]models.Device, error) {
	return repository.GetModuleGroupDevices(moduleGroupID)
}