package models

type ModuleGroup struct {
	ModuleGroupID    int     `json:"module_group_id"`
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
