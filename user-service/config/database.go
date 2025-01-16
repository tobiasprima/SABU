package config

import (
	"fmt"
	"os"
	"sabu-user-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
	User       *gorm.DB
}

var Database *DB

func InitDB() error{
	err:= godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %v", err)
	}

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
		User:       conn.Model(&models.User{}),
	}

	fmt.Println("Successfully connected to the Database")
	return nil
}