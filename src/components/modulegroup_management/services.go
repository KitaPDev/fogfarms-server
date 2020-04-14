package modulegroup_management

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/util/permission"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
)

func PopulateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	u, err := user.GetUserByUsernameFromCookie(r)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Cookie"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var jsonData []byte
	type ModuleGroupData struct {
		ModuleGroupID  int       `json:"module_group_id"`
		PlantID        int       `json:"plant_id"`
		LocationID     int       `json:"location_id"`
		TDS            float64   `json:"tds"`
		PH             float64   `json:"ph"`
		Humidity       float64   `json:"humidity"`
		OnAuto         bool      `json:"on_auto"`
		LightsOffHour  float64   `json:"lights_off_hour"`
		LightsOnHour   float64   `json:"lights_on_hour"`
		TimerLastReset time.Time `json:"timer_last_reset"`
		Permission     int
		Modules        []int
	}
	var moduleGroupMap = make(map[string]*ModuleGroupData)

	var moduleGroupIDs []int
	if u.IsAdministrator {
		moduleGroups, err := modulegroup.GetAllModuleGroups()
		if err != nil {
			msg := "Error: Failed to Get All Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		for _, mg := range moduleGroups {
			moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
			moduleGroupMap[mg.ModuleGroupLabel] = &ModuleGroupData{
				ModuleGroupID:  mg.ModuleGroupID,
				PlantID:        mg.PlantID,
				LocationID:     mg.LocationID,
				TDS:            mg.TDS,
				PH:             mg.PH,
				Humidity:       mg.Humidity,
				OnAuto:         mg.OnAuto,
				LightsOffHour:  mg.LightsOffHour,
				LightsOnHour:   mg.LightsOnHour,
				TimerLastReset: mg.TimerLastReset,
				Permission:     4,
				Modules:        []int{},
			}
		}

	} else {
		mapModuleGroupPermissions, err := permission.GetAssignedModuleGroups(u)
		if err != nil {
			msg := "Error: Failed to Get Assigned Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(mapModuleGroupPermissions)
		for mg := range mapModuleGroupPermissions {

			moduleGroupMap[mg.ModuleGroupLabel] = &ModuleGroupData{
				ModuleGroupID:  mg.ModuleGroupID,
				PlantID:        mg.PlantID,
				LocationID:     mg.LocationID,
				TDS:            mg.TDS,
				PH:             mg.PH,
				Humidity:       mg.Humidity,
				OnAuto:         mg.OnAuto,
				LightsOffHour:  mg.LightsOffHour,
				LightsOnHour:   mg.LightsOnHour,
				TimerLastReset: mg.TimerLastReset,
				Permission:     mapModuleGroupPermissions[mg],
				Modules:        []int{},
			}
			moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
		}

	}

	modules, err := module.GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs)
	if err != nil {
		msg := "Error: Failed to Get Modules By ModuleGroupIDs"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	for _, m := range modules {
		log.Println(m.ModuleLabel)
		moduleGroupMap[m.ModuleGroupLabel].Modules = append(moduleGroupMap[m.ModuleGroupLabel].Modules, m.ModuleID)
	}
	jsonData, err = json.Marshal(moduleGroupMap)
	if err != nil {
		msg := "Error: Failed to marshal JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
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
		PlantID          int       `json:"plant_id"`
		LocationID       int       `json:"location_id"`
		Humidity         float64   `json:"humidity"`
		OnAuto           bool      `json:"on_auto"`
		ModuleGroupLabel string    `json:"module_group_label"`
		LightsOffHour    float64   `json:"lights_off_hour"`
		LightsOnHour     float64   `json:"lights_on_hour"`
		TimerLastReset   time.Time `json:"timer_last_reset"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.CreateModuleGroup(input.ModuleGroupLabel, input.PlantID, input.LocationID,
		input.Humidity, input.LightsOnHour, input.LightsOffHour, input.OnAuto, input.TimerLastReset)
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
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.AssignModulesToModuleGroup(input.ModuleGroupID, input.ModuleIDs)
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
