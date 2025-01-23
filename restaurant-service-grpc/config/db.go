package config

import (
	"fmt"
	"os"
	"restaurant-service-grpc/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
	Restaurant       *gorm.DB
}

var Database *DB

func InitDB() error{
	dsn := os.Getenv("SUPABASE_URL")
	if dsn == "" {
		return fmt.Errorf("SUPABASE_URL is not set in the environment")
	}
	
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	Database = &DB{
		Connection: conn,
		Restaurant:       conn.Model(&models.Restaurant{}),
	}

	fmt.Println("Successfully connected to the Database")
	return nil
}