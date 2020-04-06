package models

type Plant struct {
	PlantID       int     `json:"plant_id"`
	Name          string  `json:"name"`
	TDS           float32 `json:"tds"`
	PH            float32 `json:"ph"`
	Lux           float32 `json:"lux"`
	LightsOnHour  float32 `json:"lights_on_hour"`
	LightsOffHour float32 `json:"lights_off_hour"`
}
