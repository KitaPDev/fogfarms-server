package modulegroup_management

import (
	"fmt"
	"net/http"
)

type DemoJson struct {
	DemoName  string
	DemoArray []string
}

func GetAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
