package user_management

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/permission"
	"github.com/KitaPDev/fogfarms-server/src/user"
	"github.com/golang/gddo/httputil/header"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user.CreateUser(w, r)
	w.WriteHeader(http.StatusOK)
}

func PopulateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	u, err := user.GetUserByUsernameFromCookie(w, r)
	if err != nil {
		msg := "Error: Failed to Get User By Username From Request"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	users, err := user.GetAllUsers()
	if err != nil {
		msg := "Error: Failed to Get All Users"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var moduleGroups []models.ModuleGroup

	if u.IsAdministrator {
		moduleGroups, err = modulegroup.GetAllModuleGroups()
		if err != nil {
			msg := "Error: Failed to Get All Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}

	} else {
		moduleGroups, err = permission.GetSupervisorModuleGroups(u)
		if err != nil {
			msg := "Error: Failed to Get Supervisor Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	var userIDs []int
	for _, u := range users {
		userIDs = append(userIDs, u.UserID)
	}

	var moduleGroupIDs []int
	for _, mg := range moduleGroups {
		moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
	}
	fmt.Printf("%+v", userIDs)
	fmt.Printf("%+v", moduleGroupIDs)
	userModuleGroupPermission, err := permission.GetUserModuleGroupPermissions(userIDs, moduleGroupIDs)
	if err != nil {
		msg := "Error: Failed to Get User ModuleGroup Permissions"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	dataJson, err := json.Marshal(userModuleGroupPermission)
	if err != nil {
		msg := "Error: json.Marshal(userModuleGroupPermission)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(w.Write(dataJson))
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

	var input Input

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
		msg := "Error: Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Printf("%+v", input)

	err = permission.AssignUserModuleGroupPermission(input.UserID, input.ModuleGroupID, input.PermissionLevel)
	if err != nil {
		msg := "Error: Failed to Assign User ModuleGroup Permission"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
