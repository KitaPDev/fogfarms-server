package main

import (
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup_management"
	"github.com/KitaPDev/fogfarms-server/src/plant_management"
	"github.com/KitaPDev/fogfarms-server/src/user_management"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	router := mux.NewRouter()

	//mux := http.NewServeMux()

	jwtAuthHandler := jwt.MakeHTTPHandler()
	router.PathPrefix("/auth").Handler(jwtAuthHandler)
	//mux.Handle("/auth", jwtAuthHandler)

	moduleGroupManagementHandler := modulegroup_management.MakeHTTPHandler()
	router.PathPrefix("/modulegroup_management").Handler(moduleGroupManagementHandler)
	//mux.Handle("/modulegroup_management", moduleGroupManagementHandler)
	//mux.Handle("/modulegroup_management/", moduleGroupManagementHandler)

	userManagementHandler := user_management.MakeHTTPHandler()
	router.PathPrefix("/user_management").Handler(userManagementHandler)
	//̧mux.Handle("/user_management", userManagementHandler)
	//mux.Handle("/user_management/", userManagementHandler)

	plantManagementHandler := plant_management.MakeHTTPHandler()
	router.PathPrefix("/plant_management").Handler(plantManagementHandler)
	//mux.Handle("/plant_management", plantManagementHandler)
	//mux.Handle("/plant_management/", plantManagementHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
