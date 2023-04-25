package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"halyklife-lib/models"
)

var DB *gorm.DB

func SetupDatabase() error {
	dsn := "host=localhost user=postgres password=darkside dbname=library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db
	DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Reader{})

	return nil
}
