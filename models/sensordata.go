package models

import "time"

type SensorData struct {
	ModuleGroupID       int       `json:"module_group_id"`
	ModuleID            int       `json:"module_id"`
	TimeStamp           time.Time `json:"timestamp"`
	TDS                 float32   `json:"tds"`
	PH                  float32   `json:"ph"`
	SolutionTemperature float32   `json:"soln_temp"`
	GrowUnitLux         []float32 `json:"grow_unit_lux"`
	GrowUnitHumidity    []float32 `json:"grow_unit_humidity"`
	GrowUnitTemperature []float32 `json:"grow_unit_temp"`
}
