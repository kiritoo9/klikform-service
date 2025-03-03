package main

import (
	"fmt"
	"klikform/src/infras/configs"
	"klikform/src/infras/databases/postgresql"
	welcomeroutes "klikform/src/interfaces/v1/routes/welcomes"

	"log"
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
	db, errDB := postgresql.OpenDB()
	if errDB != nil {
		log.Fatal("Something wrong when trying to open connection", errDB)
	}
	postgresql.Migrations(db)
	postgresql.Seeders(db)

	// regist routes
	welcomeroutes.WelcomeRoutes(mux)

	// regist swagger
	// for notes the route /swagger/doc.json is default json file loaded by swagger
	// I change it into ./docs/swagger.json for custom path
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
