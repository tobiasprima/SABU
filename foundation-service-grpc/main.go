package main

import (
	"foundation-service-grpc/config"
	"foundation-service-grpc/handlers"
	"foundation-service-grpc/proto/pb"
	"foundation-service-grpc/repository"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down servers...")

	grpcServer.GracefulStop()

	log.Println("Servers shut down gracefully")
}
