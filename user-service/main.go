package main

import (
	"log"
	"os"
	"sabu-user-service/config"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	_, err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// AutoMigrate models BUAT PAS PERTAMA JALANIN BIAR TABLE AUTO CREATE, PAS UDA CREATED DELETE GAPAPA
	// db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.ActivityLog{}) // SESUAIN SAMA MODEL KALIAN

	// routes.RegisterRoutes(e, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}