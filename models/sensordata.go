package models

import "time"

type SensorData struct {
	ModuleID            int       `json:"module_id"`
	TimeStamp           time.Time `json:"timestamp"`
	TDS                 float32   `json:"tds"`
	PH                  float32   `json:"ph"`
	SolutionTemperature float32   `json:"solution_temp"`
	GrowUnitLux         []float64 `json:"grow_unit_lux"`
	GrowUnitHumidity    []float64 `json:"grow_unit_humidity"`
	GrowUnitTemperature []float64 `json:"grow_unit_temp"`
}
