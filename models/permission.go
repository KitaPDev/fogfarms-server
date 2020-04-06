package models

type Permission struct {
	PermissionID    int `json:"permission_id"`
	PermissionLevel int `json:"permission_level"`
	UserID          int `json:"user_id"`
	ModuleGroupID   int `json:"module_group_id"`
}
