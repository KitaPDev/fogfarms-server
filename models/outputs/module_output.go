package outputs

type ModuleOutput struct {
	ModuleID         int    `json:"module_id"`
	ModuleGroupID    int    `json:"module_group_id"`
	ModuleLabel      string `json:"module_label"`
	ModuleGroupLabel string `json:"module_group_label"`
}
