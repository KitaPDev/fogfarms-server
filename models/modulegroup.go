package models

type ModuleGroup struct {
	ModuleGroupID    int     `json:"module_group_id"`
	ModuleGroupLabel string  `json:"module_group_label"`
	PlantID          int     `json:"plant_id"`
	LocationID       int     `json:"location_id"`
	TDS              float32 `json:"tds"`
	PH               float32 `json:"ph"`
	Humidity         float32 `json:"humidity"`
	LightsOnHour     float32 `json:"lights_on_hour"`
	LightsOffHour    float32 `json:"lights_off_hour"`
	OnAuto           bool    `json:"on_auto"`
}
