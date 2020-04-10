package modulegroup_management

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup/repository"
	"github.com/KitaPDev/fogfarms-server/src/util/permission"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func PopulateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	u, err := user.GetUserByUsernameFromCookie(w, r)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Request"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var jsonData []byte

	if u.IsAdministrator {
		moduleGroups, err := modulegroup.GetAllModuleGroups()
		if err != nil {
			msg := "Error: Failed to Get All Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var moduleGroupIDs []int
		for _, mg := range moduleGroups {
			moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
		}

		modules, err := module.GetModulesByModuleGroupIDs(moduleGroupIDs)
		if err != nil {
			msg := "Error: Failed to Get Modules By ModuleGroupIDs"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		mapModuleGroupModules, unassignedModules := mapModulesToModuleGroup(moduleGroups, modules)

		type Data struct {
			MapModuleGroupModules map[models.ModuleGroup][]models.Module
			UnassignedModules     []models.Module `json:"unassigned_modules"`
		}

		data := Data{
			MapModuleGroupModules: mapModuleGroupModules,
			UnassignedModules:     unassignedModules,
		}

		jsonData, err = json.Marshal(data)
		if err != nil {
			msg := "Error: Failed to marshal JSON"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}

	} else {
		mapModuleGroupPermissions, err := permission.GetAssignedModuleGroups(u)
		if err != nil {
			msg := "Error: Failed to Get Assigned Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var moduleGroups []models.ModuleGroup
		var moduleGroupIDs []int
		for mg := range mapModuleGroupPermissions {
			moduleGroups = append(moduleGroups, mg)
			moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
		}

		modules, err := module.GetModulesByModuleGroupIDs(moduleGroupIDs)
		if err != nil {
			msg := "Error: Failed to Get Modules By ModuleGroupIDs"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		mapModuleGroupModules, unassignedModules := mapModulesToModuleGroup(moduleGroups, modules)

		type Data struct {
			MapModuleGroupModules    map[models.ModuleGroup][]models.Module
			MapModuleGroupPermission map[models.ModuleGroup]int
			UnassignedModules        []models.Module `json:"unassigned_modules"`
		}

		data := Data{
			MapModuleGroupModules:    mapModuleGroupModules,
			MapModuleGroupPermission: mapModuleGroupPermissions,
			UnassignedModules:        unassignedModules,
		}

		jsonData, err = json.Marshal(data)
		if err != nil {
			msg := "Error: Failed to marshal JSON"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CreateModuleGroup(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	type Input struct {
		PlantID          int     `json:"plant_id"`
		LocationID       int     `json:"location_id"`
		TDS              float32 `json:"tds"`
		PH               float32 `json:"ph"`
		Humidity         float32 `json:"humidity"`
		OnAuto           bool    `json:"on_auto"`
		ModuleGroupLabel string  `json:"module_group_label"`
		LightsOffHour    float32 `json:"lights_off_hour"`
		LightsOnHour     float32 `json:"lights_on_hour"`
	}

	input := Input{}
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

	err = repository.CreateModuleGroup(input.ModuleGroupLabel, input.PlantID, input.LocationID,
		input.Humidity, input.LightsOnHour, input.LightsOffHour, input.OnAuto)
	if err != nil {
		msg := "Error: Failed to Create Module Group"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Create ModuleGroup; Successful"))
}

func AssignModuleToModuleGroup(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	type Input struct {
		ModuleGroupID int   `json:"module_group_ids"`
		ModuleIDs     []int `json:"module_ids"`
	}

	input := Input{}
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

	err = repository.AssignModulesToModuleGroup(input.ModuleGroupID, input.ModuleIDs)
	if err != nil {
		msg := "Error: Failed to Assign Modules To ModuleGroup"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Assign Module To ModuleGroup; Successful"))
}

func mapModulesToModuleGroup(moduleGroups []models.ModuleGroup,
	modules []models.Module) (map[models.ModuleGroup][]models.Module, []models.Module) {

	mapModuleGroupModules := make(map[models.ModuleGroup][]models.Module)
	var assignedModules []models.Module
	var unassignedModules []models.Module

	for _, mg := range moduleGroups {
		mapModuleGroupModules[mg] = make([]models.Module, 0)

		for _, m := range modules {
			if m.ModuleGroupID == mg.ModuleGroupID {
				mapModuleGroupModules[mg] = append(mapModuleGroupModules[mg], m)
				assignedModules = append(assignedModules, m)
			}
		}
	}

	for _, m1 := range modules {
		found := false

		for _, m2 := range assignedModules {
			if m2 == m1 {
				found = true
				break
			}
		}

		if !found {
			unassignedModules = append(unassignedModules, m1)
		}
	}

	return mapModuleGroupModules, unassignedModules
}
