package permission

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/permission/repository"
	"github.com/KitaPDev/fogfarms-server/src/user"
)

func GetUserModuleGroupPermissions(userIDs []int, moduleGroupIDs []int) (map[string]map[string]int, error) {
	if len(userIDs) == 0 || len(moduleGroupIDs) == 0 {
		return make(map[string]map[string]int), nil
	}

	userModuleGroupPermissions := make(map[string]map[string]int)

	permissions := repository.GetAllPermissions()
	users, err := user.GetUsersByID(userIDs)
	if err != nil {
		return make(map[string]map[string]int), err
	}

	moduleGroups := modulegroup.GetModuleGroupsByID(moduleGroupIDs)

	fGetUsername := func(userID int) string {
		for _, uTemp := range users {
			if uTemp.UserID == userID {
				return uTemp.Username
			}
		}
		return ""
	}

	fGetModuleGroupLabel := func(moduleGroupID int) string {
		for _, mg := range moduleGroups {
			if mg.ModuleGroupID == moduleGroupID {
				return mg.ModuleGroupLabel
			}
		}
		return ""
	}

	fGetPermission := func(userID int, moduleGroupID int) int {
		for _, p := range permissions {
			if p.UserID == userID && p.ModuleGroupID == moduleGroupID {
				return p.PermissionLevel
			}
		}
		return 0
	}

	for _, uid := range userIDs {
		userModuleGroupPermissions[fGetUsername(uid)] = make(map[string]int)

		for mgid := range moduleGroupIDs {
			userModuleGroupPermissions[fGetUsername(uid)][fGetModuleGroupLabel(mgid)] =
				fGetPermission(uid, mgid)
		}

	}

	return userModuleGroupPermissions, nil
}

func AssignUserModuleGroupPermission(userID int, moduleGroupID int, permissionLevel int) error {
	return repository.AssignUserModuleGroupPermission(userID, moduleGroupID, permissionLevel)
}

func GetSupervisorModuleGroups(user *models.User) []models.ModuleGroup {
	return repository.GetSupervisorModuleGroups(user.UserID)
}