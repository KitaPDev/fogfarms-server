package models

type Device struct {
	DeviceID int `json:"device_id"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}