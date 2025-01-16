package main

import (
	"log"
	"os"
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/routes"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	routes.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}