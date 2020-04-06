package permission

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/permission/repository"
	"github.com/KitaPDev/fogfarms-server/src/user"
)

func GetUserModuleGroupPermissions(userIDs []int, moduleGroupIDs []int) map[string]map[string]int {
	if len(userIDs) == 0 || len(moduleGroupIDs) == 0 {
		return make(map[string]map[string]int)
	}

	permissions := repository.GetAllPermissions()
	userModuleGroupPermissions := make(map[string]map[string]int)
	users := user.GetUsersByID(userIDs)
	moduleGroups := modulegroup.GetModuleGroupsByID()

	for _, uid := range userIDs {
		u := user.GetUserByID(uid)
		userModuleGroupPermissions[]

		for mgid := range moduleGroupIDs {


		}

	}

	return userModuleGroupPermissions
}

func AssignUserToModuleGroup(username string, moduleGroupID int, permissionLevel int) {
	repository.AssignUserToModuleGroup(username, moduleGroupID, permissionLevel)
}

func GetSupervisorModuleGroups(user *models.User) []models.ModuleGroup {
	return repository.GetSupervisorModuleGroups(user.UserID)
}