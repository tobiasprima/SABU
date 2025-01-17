package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/handlers"
	"sabu-restaurant-service/proto/pb"
	"sabu-restaurant-service/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	err = config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	err = config.InitMongo()
	if err != nil {
		log.Fatalf("Mongodb connection failed: %v", err)
	}

	grpcServer := grpc.NewServer()

	restaurantHandler := handlers.NewRestaurantGRpcHandler()

	// Register the gRPC handler with the gRPC server
	pb.RegisterRestaurantServiceServer(grpcServer, restaurantHandler)


	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on port 50051: %v", err)
		}

		log.Println("gRPC server running on port 50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	
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

    grpcServer.GracefulStop()
    if err := e.Shutdown(context.Background()); err != nil {
        log.Fatalf("Failed to shut down REST server: %v", err)
    }

    log.Println("Servers shut down gracefully")
}