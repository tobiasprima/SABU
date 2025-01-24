package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("failed to load .env file: %v", err)
	// }

	err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	err = config.InitMongo()
	if err != nil {
		log.Fatalf("Mongodb connection failed: %v", err)
	}

	e := echo.New()
	routes.RegisterRoutes(e)
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8081"
		}

		log.Printf("Starting server on port %s", port)
		e.Logger.Fatal(e.Start(":" + port))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down servers...")

	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Failed to shut down REST server: %v", err)
	}

	log.Println("Servers shut down gracefully")
}
