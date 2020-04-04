package models

type Level string

const (
	Supervisor Level = "Supervisor"
	Control    Level = "Control"
	Monitor    Level = "Monitor"
)

type Permission struct {
	PermissionID  string `json:"permission_id"`
	Level         Level  `json:"level"`
	UserID        string `json:"user_id"`
	ModuleGroupID string `json:"module_group_id"`
}
