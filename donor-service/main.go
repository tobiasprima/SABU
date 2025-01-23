package main

import (
	"context"
	"crypto/tls"
	"donor-service/config"
	"donor-service/handlers"
	"donor-service/proto/pb"
	"donor-service/repository"
	"donor-service/routes"
	"donor-service/service"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	foundationgRPCconn, err := grpc.Dial(os.Getenv("FOUNDATION_SERVICE"), opts...)
	if err != nil {
		log.Fatalf("Failed to connect to Foundation service: %v", err)
	}
	defer foundationgRPCconn.Close()

	creds2 := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts2 := []grpc.DialOption{
		grpc.WithTransportCredentials(creds2),
	}

	restaurantgRPCconn, err := grpc.Dial(os.Getenv("RESTAURANT_SERVICE"), opts2...)
	if err != nil {
		log.Fatalf("Failed to connect to Restaurant service: %v", err)
	}
	defer restaurantgRPCconn.Close()

	foundationClient := pb.NewFoundationServiceClient(foundationgRPCconn)
	restaurantClient := pb.NewRestaurantServiceClient(restaurantgRPCconn)

	donorRepository := repository.NewDonorRepositoryImpl(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	paymentService := service.NewPaymentService(os.Getenv("XENDIT_API_KEY"))
	donorHandler := handlers.NewDonorHandlerImpl(donorRepository, paymentService, foundationClient, restaurantClient)

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

	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Failed to shut down REST server: %v", err)
	}

	log.Println("Servers shut down gracefully")
}
