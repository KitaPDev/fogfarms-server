package models

type NutrientUnit struct {
	NutrientUnitID int `json:"nutrient_unit_id"`
	ModuleID       int `json:"module_id"`
	ModuleGroupID  int `json:"module_group_id"`
	PHUpUnitID     int `json:"ph_up_unit_id"`
	PHDownUnitID   int `json:"ph_down_unit_id"`
	NutrientID     int `json:"nutrient_id"`
}
