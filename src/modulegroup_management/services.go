package modulegroup_management

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup/repository"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func PopulateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	//if !jwt.AuthenticateUserToken(w, r) {
	//msg := "Unauthorized"
	//http.Error(w, msg, http.StatusUnauthorized)
	//return
	//}
	//
	//u, err := user.GetUserByUsernameFromCookie(w, r)
	//if err != nil {
	//	msg := "Error: Failed to Get User By Username From Request"
	//	http.Error(w, msg, http.StatusInternalServerError)
	//	log.Println(err)
	//	return
	//}
	//
	//if u.IsAdministrator {
	//	moduleGroups, err := modulegroup.GetAllModuleGroups()
	//	if err != nil {
	//		msg := "Error: Failed to Get All Module Groups"
	//		http.Error(w, msg, http.StatusInternalServerError)
	//		log.Println(err)
	//		return
	//	}
	//
	//
	//
	//} else {
	//	mapModuleGroupPermissions, err := permission.GetAssignedModuleGroups(u)
	//	if err != nil {
	//		msg := "Error: Failed to Get Assigned Module Groups"
	//		http.Error(w, msg, http.StatusInternalServerError)
	//		log.Println(err)
	//		return
	//	}
	//
	//
	//
	//}
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