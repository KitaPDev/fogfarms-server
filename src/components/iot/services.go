package iot

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata"
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Token          string               `json:"token"`
		SensorData     map[string][]float64 `json:"sensor_data"`
		Mixers         []bool               `json:"mixers"`
		SolenoidValves []bool               `json:"solenoid_valves"`
		LEDs           []bool               `json:"leds"`
		Foggers        []bool               `json:"foggers"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	moduleID, err := module.GetModuleIDByToken(input.Token)
	if err != nil {
		msg := "Error: Failed to Get ModuleID By Token"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	onAuto, err := modulegroup.GetOnAutoByModuleID(moduleID)
	if err != nil {
		msg := "Error: Failed to Get OnAuto By ModuleID"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	tds := input.SensorData["tds"]
	ph := input.SensorData["ph"]
	solutionTemp := input.SensorData["solution_temp"]
	lux := input.SensorData["grow_unit_lux"]
	humidity := input.SensorData["grow_unit_humidity"]
	temperature := input.SensorData["grow_unit_temp"]

	err = sensordata.RecordSensorData(moduleID, tds, ph, solutionTemp, lux, humidity, temperature)
	if err != nil {
		msg := "Error: Failed to Record Sensor Data"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if onAuto {
		err = module.UpdateDeviceStatus(moduleID, input.Mixers, input.SolenoidValves, input.LEDs,
			input.Foggers)
		if err != nil {
			msg := "Error: Failed to Record Device Status"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successful"))

	} else {
		fogger, led, mixer, solenoidValve, err := module.GetDeviceStatus(moduleID)
		if err != nil {
			msg := "Error: Failed to Get Device Status"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		type DeviceStatus struct {
			Fogger        []bool `json:"fogger"`
			LED           []bool `json:"led"`
			Mixer         []bool `json:"mixer"`
			SolenoidValve []bool `json:"solenoid_valve"`
		}

		deviceStatus := DeviceStatus{
			Fogger:        fogger,
			LED:           led,
			Mixer:         mixer,
			SolenoidValve: solenoidValve,
		}

		jsonData, err := json.Marshal(deviceStatus)
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
}