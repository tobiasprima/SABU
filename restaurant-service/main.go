package main

import (
	"log"
	"os"
	"sabu-restaurant-service/config"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// AutoMigrate models BUAT PAS PERTAMA JALANIN BIAR TABLE AUTO CREATE, PAS UDA CREATED DELETE GAPAPA
	// db.AutoMigrate(&models.User{}) // SESUAIN SAMA MODEL KALIAN

	// routes.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}