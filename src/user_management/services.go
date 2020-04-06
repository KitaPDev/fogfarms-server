package user_management

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/permission"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user.CreateUser(w, r)
	w.WriteHeader(http.StatusOK)
}

func PopulateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUser(w, r) {
		return
	}

	u := user.GetUserByUsernameFromRequest(w, r)

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

	userModuleGroupPermission := permission.GetUserModuleGroupPermissions(userIDs, moduleGroupIDs)

	dataJson, err := json.Marshal(userModuleGroupPermission)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	log.Fatal(w.Write(dataJson))
}

func AssignUserModuleGroupPermission(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUser(w, r) {
		return
	}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		panic(err)
	}
}