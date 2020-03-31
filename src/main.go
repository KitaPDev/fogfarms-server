package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/plant_management"
	"github.com/KitaPDev/fogfarms-server/src/user_management"
	"github.com/ddfsdd/fogfarms-server/src/modulegroup_management"
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
	mux := http.NewServeMux()

	jwtAuthHandler := jwt.MakeHTTPHandler()
	mux.Handle("/auth", jwtAuthHandler)

	moduleGroupManagementHandler := modulegroup_management.MakeHTTPHandler()
	mux.Handle("/modulegroup_management", moduleGroupManagementHandler)
	mux.Handle("/modulegroup_management/js", moduleGroupManagementHandler)
	userManagementHandler := user_management.MakeHTTPHandler()
	mux.Handle("/user_management", userManagementHandler)

	plantManagementHandler := plant_management.MakeHTTPHandler()
	mux.Handle("/plant_management", plantManagementHandler)

	return http.ListenAndServe(getPort(), mux)
}
