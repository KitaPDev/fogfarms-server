package models

type GrowUnit struct {
	GrowUnitID    int `json:"grow_unit_id"`
	ModuleID      int `json:"module_id"`
	ModuleGroupID int `json:"module_group_id"`
	Capacity      int `json:"capacity"`
}