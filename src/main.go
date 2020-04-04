package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup_management"
	"github.com/KitaPDev/fogfarms-server/src/plant_management"
	"github.com/KitaPDev/fogfarms-server/src/test"
	"github.com/KitaPDev/fogfarms-server/src/user_management"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9090"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}
func run() error {
	router := mux.NewRouter()

	jwtAuthHandler := jwt.MakeHTTPHandler()
	router.PathPrefix("/auth").Handler(jwtAuthHandler)

	moduleGroupManagementHandler := modulegroup_management.MakeHTTPHandler()
	router.PathPrefix("/modulegroup_management").Handler(moduleGroupManagementHandler)

	userManagementHandler := user_management.MakeHTTPHandler()
	router.PathPrefix("/user_management").Handler(userManagementHandler)

	plantManagementHandler := plant_management.MakeHTTPHandler()
	router.PathPrefix("/plant_management").Handler(plantManagementHandler)

	return http.ListenAndServe(getPort(), router)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9090"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}