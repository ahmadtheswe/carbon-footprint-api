package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	if AppConfig == nil {
		log.Fatal("Config not loaded. Call LoadConfig() first.")
	}

	dsn := AppConfig.DBStringConnection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	return db
}
