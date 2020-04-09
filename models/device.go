package models

type Device struct {
	DeviceID int    `json:"device_id"`
	IsOn     bool   `json:"IsOn"`
	ModuleID int    `json:"module_id"`
	Label    string `json:"label"`
}
