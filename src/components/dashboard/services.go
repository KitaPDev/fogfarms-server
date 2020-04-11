package dashboard

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/util/device"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata"
	"log"
	"net/http"
)

func PopulateDashboard(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	moduleGroup := models.ModuleGroup{}

	success := jsonhandler.DecodeJsonFromBody(w, r, &moduleGroup)
	if !success {
		return
	}

	sensorData, err := sensordata.GetLatestSensorData(moduleGroup.ModuleGroupID)
	if err != nil {
		msg := "Error: Failed to Get Latest Sensor Data"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	devices, err := device.GetModuleGroupDevices(moduleGroup.ModuleGroupID)
	if err != nil {
		msg := "Error: Failed to Get ModuleGroup Devices"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	type Data struct {
		SensorData []models.SensorData
		Devices    []models.Device
	}

	data := Data{
		SensorData: sensorData,
		Devices:    devices,
	}

	jsonData, err := json.Marshal(data)
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

func ToggleDevice(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	var deviceID int

	success := jsonhandler.DecodeJsonFromBody(w, r, &deviceID)
	if !success {
		return
	}



}