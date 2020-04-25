package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KitaPDev/fogfarms-server/src/components/iot"
	"github.com/KitaPDev/fogfarms-server/src/components/module_management"
	"github.com/KitaPDev/fogfarms-server/src/test"
	"github.com/rs/cors"

	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/components/dashboard"
	"github.com/KitaPDev/fogfarms-server/src/components/modulegroup_management"
	"github.com/KitaPDev/fogfarms-server/src/components/plant_management"
	"github.com/KitaPDev/fogfarms-server/src/components/user_management"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	router := mux.NewRouter()

	jwtAuthHandler := jwt.MakeHTTPHandler()
	router.PathPrefix("/auth").Handler(jwtAuthHandler)

	moduleGroupManagementHandler := modulegroup_management.MakeHTTPHandler()
	router.PathPrefix("/modulegroup_management").Handler(moduleGroupManagementHandler)

	moduleManagementHandler := module_management.MakeHTTPHandler()
	router.PathPrefix("/module_management").Handler(moduleManagementHandler)

	userManagementHandler := user_management.MakeHTTPHandler()
	router.PathPrefix("/user_management").Handler(userManagementHandler)

	plantManagementHandler := plant_management.MakeHTTPHandler()
	router.PathPrefix("/plant_management").Handler(plantManagementHandler)

	dashBoardHandler := dashboard.MakeHTTPHandler()
	router.PathPrefix("/dashboard").Handler(dashBoardHandler)

	iotHandler := iot.MakeHTTPHandler()
	router.PathPrefix("/iot").Handler(iotHandler)

	testHandler := test.MakeHTTPHandler()
	router.PathPrefix("/test").Handler(testHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://localhost:3000", "https://25.22.245.97:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}).Handler(router)
	return http.ListenAndServe(getPort(), handler)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	fmt.Println("Server is running on port: " + port)
	return ":" + port
}
