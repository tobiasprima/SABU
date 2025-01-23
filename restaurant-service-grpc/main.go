package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"restaurant-service-grpc/config"
	"restaurant-service-grpc/handlers"
	"restaurant-service-grpc/proto/pb"

	"google.golang.org/grpc"
)


func main() {
	err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	err = config.InitMongo()
	if err != nil {
		log.Fatalf("Mongodb connection failed: %v", err)
	}

	grpcServer := grpc.NewServer()

	restaurantHandler := handlers.NewRestaurantGRpcHandler()

	pb.RegisterRestaurantServiceServer(grpcServer, restaurantHandler)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

	quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    log.Println("Shutting down servers...")

    grpcServer.GracefulStop()

    log.Println("Servers shut down gracefully")

}