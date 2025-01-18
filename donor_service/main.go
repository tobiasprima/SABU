package main

import (
	"context"
	"donor-service/config"
	"donor-service/handlers"
	"donor-service/proto/pb"
	"donor-service/repository"
	"donor-service/routes"
	"donor-service/service"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	paymentService := service.NewPaymentService(os.Getenv("XENDIT_API_KEY"))
	donorHandler := handlers.NewDonorHandlerImpl(donorRepository, paymentService)

	routes.RegisterRoutes(e, *donorHandler)

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8082"
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
