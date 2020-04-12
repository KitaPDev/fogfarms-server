package models

type Module struct {
	ModuleID      int    `json:"module_id"`
	ModuleGroupID int    `json:"module_group_id"`
	Token         string `json:"token"`
	ModuleLabel   string `json:"module_label`
}

type ModuleOutput struct {
	ModuleID         int
	ModuleGroupID    int
	ModuleLabel      string
	ModuleGroupLabel string
}
