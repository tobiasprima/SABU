package main

import (
	"donor-service/config"
	"donor-service/handlers"
	"donor-service/repository"
	"donor-service/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	donorRepository := repository.NewDonorRepositoryImpl(db)
	donorHandler := handlers.NewDonorHandlerImpl(donorRepository)

	routes.RegisterRoutes(e, *donorHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
