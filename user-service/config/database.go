package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	err:= godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	dsn := os.Getenv("SUPABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("SUPABASE_URL is not set in the environment")
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	fmt.Println("Successfully connected to the Database")
	return db, nil
}