package models

type GrowUnit struct {
	GrowUnitID    int `json:"grow_unit_id"`
	ModuleID      int `json:"module_id"`
	Capacity      int `json:"capacity"`
}
