package main

import (
	"fmt"
	"klikform/src/infras/configs"
	"klikform/src/infras/databases/postgresql"
	authroutes "klikform/src/interfaces/v1/routes/auths"
	masterroutes "klikform/src/interfaces/v1/routes/masters"
	welcomeroutes "klikform/src/interfaces/v1/routes/welcomes"

	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title KlikForm Service
// @version 0.1
// @description KlikForm Service API Documentation
func main() {
	mux := http.NewServeMux()

	// load necessary function that needs to call when app init
	// environments - database connections - migrations - seeders
	conf := configs.LoadConfig()
	postgresql.OpenDB() // run database connection

	// uncomment this if you want to run migrations and seeders
	// migrations.Migrations()
	// seeders.Seeders()

	// regist initiate routes
	welcomeroutes.WelcomeRoutes(mux)
	authroutes.AuthRoutes(mux)

	// regist master routes
	masterroutes.RoleRoutes(mux)
	masterroutes.UserRoutes(mux)

	// regist swagger
	// for notes the route /swagger/doc.json is default json file loaded by swagger
	// I change it into ./docs/swagger.json for custom path
	// @title        KlikForm API
	// @version      1.0
	// @description  Klikform API Service

	// @securityDefinitions.apikey BearerAuth
	// @in           header
	// @name         Authorization

	// @security     BearerAuth
	// @accept       json
	// @produce      json
	// @param Authorization header string true "Bearer {token}"
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	// run the app
	port := ":" + conf.APP_PORT
	fmt.Println("Server is running on " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
