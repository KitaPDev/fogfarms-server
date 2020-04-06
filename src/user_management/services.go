package user_management

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/permission"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	repository.GetAllUsers()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user.CreateUser(w, r)
}

func PopulateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUser(w, r) {
		return
	}

	u := user.GetUserFromRequest(w, r)

	type moduleGroupPermission struct {
		moduleGroupLabel string
		permissionLevel int
	}

	var users []models.User
	var moduleGroups []models.ModuleGroup

	if u.IsAdministrator {
		users = user.GetAllUsers()
		moduleGroups = modulegroup.GetAllModuleGroups()

	} else {
		users = user.GetAllUsers()
		moduleGroups = permission.GetSupervisorModuleGroups(u)
	}

	var userIDs []int
	for _, u := range users {
		userIDs = append(userIDs, u.UserID)
	}

	var moduleGroupIDs []int
	for _, mg := range moduleGroups {
		moduleGroupIDs = append (moduleGroupIDs, mg.ModuleGroupID)
	}


}