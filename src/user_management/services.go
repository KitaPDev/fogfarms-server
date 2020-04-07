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
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	u, err := user.GetUserByUsernameFromRequest(w, r)
	if err != nil {
		msg := "Error: user.GetUserByUsernameFromRequest"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	users, err := user.GetAllUsers()
	if err != nil {
		msg := "Error: user.GetAllUsers()"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	var moduleGroups []models.ModuleGroup

	if u.IsAdministrator {
		moduleGroups, err = modulegroup.GetAllModuleGroups()
		if err != nil {
			msg := "Error: moduleGroup.GetAllModuleGroups()"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

	} else {
		moduleGroups = permission.GetSupervisorModuleGroups(u)
	}

	var userIDs []int
	for _, u := range users {
		userIDs = append(userIDs, u.UserID)
	}

	var moduleGroupIDs []int
	for _, mg := range moduleGroups {
		moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
	}

	userModuleGroupPermission, err := permission.GetUserModuleGroupPermissions(userIDs, moduleGroupIDs)
	if err != nil {
		msg := "Error: permission.GetUserModuleGroupPermission(userIDs, moduleGroupIDs)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	dataJson, err := json.Marshal(userModuleGroupPermission)
	if err != nil {
		msg := "Error: json.Marshal(userModuleGroupPermission)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Fatal(w.Write(dataJson))
}

func AssignUserModuleGroupPermission(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		UserID          int `json:"user_id"`
		ModuleGroupID   int `json:"module_group_id"`
		PermissionLevel int `json:"permission_level"`
	}

	var input *Input

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		msg := "Error: json.NewDecoder(r.Body).Decode(&input)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	err = permission.AssignUserModuleGroupPermission(input.UserID, input.ModuleGroupID, input.PermissionLevel)
	if err != nil {
		msg := "Error: permission.AssignUserModuleGroupPermission(input.UserID, input.ModuleGroupID, input.PermissionLevel)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}