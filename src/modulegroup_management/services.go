package modulegroup_management

import (
	"net/http"
)

func PopulateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	//if !jwt.AuthenticateUserToken(w, r) {
	//	return
	//}
	//
	//u, err := user.GetUserByUsernameFromCookie(w, r)
	//if err != nil {
	//	msg := "Error: Failed to Get User By Username From Request"
	//	http.Error(w, msg, http.StatusInternalServerError)
	//	log.Println(err)
	//	return
	//}
	//
	//if u.IsAdministrator {
	//	moduleGroups, err := modulegroup.GetAllModuleGroups()
	//	if err != nil {
	//		msg := "Error: Failed to Get All Module Groups"
	//		http.Error(w, msg, http.StatusInternalServerError)
	//		log.Println(err)
	//		return
	//	}
	//
	//
	//
	//} else {
	//	mapModuleGroupPermissions, err := permission.GetAssignedModuleGroups(u)
	//	if err != nil {
	//		msg := "Error: Failed to Get Assigned Module Groups"
	//		http.Error(w, msg, http.StatusInternalServerError)
	//		log.Println(err)
	//		return
	//	}
	//
	//
	//
	//}
}