package main

import (
	"crypto/tls"
	"log"
	"os"
	"sabu-user-service/config"
	"sabu-user-service/proto/pb"
	"sabu-user-service/routes"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	e := echo.New()

	err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	restaurantgRPCconn, err := grpc.Dial(os.Getenv("RESTAURANT_SERVICE"), opts...)
	if err != nil {
		log.Fatalf("Failed to connect to Restaurant service: %v", err)
	}
	defer restaurantgRPCconn.Close()

	creds2 := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts2 := []grpc.DialOption{
		grpc.WithTransportCredentials(creds2),
	}

	donorgRPCConn, err := grpc.Dial(os.Getenv("DONOR_SERVICE"), opts2...)
	if err != nil {
		log.Fatalf("Failed to connect to Donor service: %v", err)
	}
	defer donorgRPCConn.Close()

	creds3 := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts3 := []grpc.DialOption{
		grpc.WithTransportCredentials(creds3),
	}

	foundationgRPCConn, err := grpc.Dial(os.Getenv("FOUNDATION_SERVICE"), opts3...)
	if err != nil {
		log.Fatalf("Failed to connect to Foundation service: %v", err)
	}
	defer foundationgRPCConn.Close()

	restaurantClient := pb.NewRestaurantServiceClient(restaurantgRPCconn)
	donorClient := pb.NewDonorServiceClient(donorgRPCConn)
	foundationClient := pb.NewFoundationServiceClient(foundationgRPCConn)

	routes.RegisterRoutes(e, restaurantClient, donorClient, foundationClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
