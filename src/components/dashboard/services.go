package dashboard

import (
	"encoding/json"
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func PopulateDashboard(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	u, err := user.GetUserByUsernameFromCookie(w, r)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Request"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	moduleGroup := models.ModuleGroup{}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err = json.NewDecoder(r.Body).Decode(&moduleGroup)
	if err != nil {
		msg := "Error: Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}



	if u.IsAdministrator {


	} else {


	}


}