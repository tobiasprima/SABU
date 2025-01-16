package main

import (
	"log"
	"os"
	"sabu-user-service/config"
	"sabu-user-service/proto/pb"
	"sabu-user-service/routes"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main(){
	e := echo.New()

	err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Restaurant service: %v", err)
	}
	defer grpcConn.Close()

	restaurantClient := pb.NewRestaurantServiceClient(grpcConn)

	routes.RegisterRoutes(e, restaurantClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}