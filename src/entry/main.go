package main

import (
	"fmt"
	"klikform/src/infras/configs"
	"klikform/src/interfaces/v1/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// load necessary function that needs to call when app init
	// environments - database connections - swaggers
	conf := configs.LoadConfig()

	// regist routes
	routes.WelcomeRoutes(mux)

	port := ":" + conf.APP_PORT
	fmt.Println("Server is running on " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
