package main

import (
	"context"
	"foundation-service/config"
	"foundation-service/handlers"
	"foundation-service/proto/pb"
	"foundation-service/repository"
	"foundation-service/routes"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
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

	grpcServer := grpc.NewServer()

	foundationRepository := repository.NewFoundationRepositoryImpl(db)
	foundationGrpcHandler := handlers.NewFoundationGrpcHandlerImpl(foundationRepository)

	pb.RegisterFoundationServiceServer(grpcServer, foundationGrpcHandler)

	go func() {
		listener, err := net.Listen("tcp", ":50053")
		if err != nil {
			log.Fatalf("Failed to listen on port 50053: %v", err)
		}

		log.Println("gRPC server running on port 50053")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	grpcServer.GracefulStop()
	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Failed to shut down REST server: %v", err)
	}

	log.Println("Servers shut down gracefully")
}
