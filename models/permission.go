package models

type Permission struct {
	PermissionID    string `json:"permission_id"`
	PermissionLevel int    `json:"permission_level"`
	UserID          string `json:"user_id"`
	ModuleGroupID   string `json:"module_group_id"`
}
