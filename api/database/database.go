package database

import (
	"log"
	"os"

	"daveslist-emdpcv/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("daveslist.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	err = DB.Migrator().DropTable(&models.Picture{}, &models.Reply{}, &models.PrivateMessage{}, &models.Listing{}, &models.Category{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
		os.Exit(1)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Listing{},
		&models.Picture{},
		&models.Reply{},
		&models.PrivateMessage{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		os.Exit(1)
	}
	log.Println("Database migration completed successfully")
}
