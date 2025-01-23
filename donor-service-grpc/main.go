package main

import (
	"donor-service-grpc/config"
	"donor-service-grpc/handlers"
	"donor-service-grpc/proto/pb"
	"donor-service-grpc/repository"
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

	donorRepository := repository.NewDonorRepositoryImpl(db)
	donorGrpcHandler := handlers.NewDonorGrpcHandlerImpl(donorRepository)

	pb.RegisterDonorServiceServer(grpcServer, donorGrpcHandler)

	go func() {
		listener, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("Failed to listen on port 50052: %v", err)
		}

		log.Println("gRPC server running on port 50052")
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
