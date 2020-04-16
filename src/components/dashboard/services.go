package dashboard

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/device"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata"
)

func PopulateDashboard(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}
	type Input struct {
		ModuleGroupID int `json:"modulegroupid"`
	}
	var input Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}
	log.Println("this is input: ", input)
	sensorData, err := sensordata.GetLatestSensorData(input.ModuleGroupID)
	if err != nil {
		msg := "Error: Failed to Get Latest Sensor Data"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	devices, err := device.GetModuleGroupDevices(input.ModuleGroupID)
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
		return
	}
	type Input struct {
		DeviceID int `json:"device_id"`
	}
	var input Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := device.ToggleDevice(input.DeviceID)
	if err != nil {
		msg := "Error: Failed to Toggle Device"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func ToggleAuto(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleGroupID int `json:"module_group_id"`
	}
	var input Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.ToggleAuto(input.ModuleGroupID)
	if err != nil {
		msg := "Error: Failed to Toggle Auto"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func SetEnvironmentParameters(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleGroupID int     `json:"module_group_id"`
		TDS           float64 `json:"tds"`
		PH            float64 `json:"ph"`
		Humidity      float64 `json:"humidity"`
		LightsOnHour  float64 `json:"lights_on_hour"`
		LightsOffHour float64 `json:"lights_off_hour"`
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
	w.Write([]byte("Successful"))
}

func ResetTimer(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}
	type Input struct {
		ModuleGroupID int `json:"module_group_id"`
	}
	var input Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.ResetTimer(input.ModuleGroupID)
	if err != nil {
		msg := "Error: Failed to Reset Timer"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}
