package database

import (
	"log"
	"user-management/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=manaf password=12345 dbname=user port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Printf("connected to database: %v", db.Name())
	db.AutoMigrate(&models.User{})
	return db
}


