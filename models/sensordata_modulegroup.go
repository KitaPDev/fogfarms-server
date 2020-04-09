package models

import "time"

type SensorDataModuleGroup struct {
	ModuleGroupID int       `json:"module_group_id"`
	Timestamp     time.Time `json:"timestamp"`
	Humidity      float32   `json:"humidity"`
	Temperature   float32   `json:"temperature"`
}
