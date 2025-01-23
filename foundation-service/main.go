package main

import (
	"context"
	"foundation-service/config"
	"foundation-service/handlers"
	"foundation-service/repository"
	"foundation-service/routes"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	foundationRepository := repository.NewFoundationRepositoryImpl(db)
	foundationHandler := handlers.NewFoundationHandlerImpl(foundationRepository)

	routes.RegisterRoutes(e, foundationHandler)

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8083"
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
