package modulegroup_management

import (
	"fmt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup/repository"
	"net/http"
)

func GetAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	moduleGroups, err := repository.GetAllModuleGroups()
	if err != nil {
		msg := "Error: Failed to Get All Module Groups"

	}
}