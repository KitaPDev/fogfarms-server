package dashboard

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/util/device"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
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

	var moduleGroupID int

	success := jsonhandler.DecodeJsonFromBody(w, r, &moduleGroupID)
	if !success {
		return
	}

	sensorData, err := sensordata.GetLatestSensorData(moduleGroupID)
	if err != nil {
		msg := "Error: Failed to Get Latest Sensor Data"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	devices, err := device.GetModuleGroupDevices(moduleGroupID)
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

	err := device.ToggleDevice(deviceID)
	if err != nil {
		msg := "Error: Failed to Toggle Device"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Toggle Device; Successful"))
}

func ToggleAuto(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	var moduleGroupID int

	success := jsonhandler.DecodeJsonFromBody(w, r, &moduleGroupID)
	if !success {
		return
	}

	err := modulegroup.ToggleAuto(moduleGroupID)
	if err != nil {
		msg := "Error: Failed to Toggle Auto"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Toggle Auto; Successful"))
}

func SetEnvironmentParameters(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	type Input struct {
		ModuleGroupID int     `json:"module_group_id"`
		TDS           float32 `json:"tds"`
		PH            float32 `json:"ph"`
		Humidity      float32 `json:"humidity"`
		LightsOnHour  float32 `json:"lights_on_hour"`
		LightsOffHour float32 `json:"lights_off_hour"`
	}

	input := Input{}

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.SetEnvironmentParameters(input.ModuleGroupID, input.Humidity, input.PH,
		input.TDS, input.LightsOnHour, input.LightsOffHour)
	if err != nil {
		msg := "Error: Failed to Set Environment Parameters"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Set Environment Parameters; Successful"))
}

func ResetTimer(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	var moduleGroupID int

	success := jsonhandler.DecodeJsonFromBody(w, r, &moduleGroupID)
	if !success {
		return
	}

	err := modulegroup.ResetTimer(moduleGroupID)
	if err != nil {
		msg := "Error: Failed to Reset Timer"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation: Reset Timer; Successful"))
}