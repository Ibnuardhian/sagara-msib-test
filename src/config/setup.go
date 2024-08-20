package config

import (
	"os"
	"test_sagara/src/models"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	connection := os.Getenv("DB_URL")
	if connection == "" {
		log.Fatal().Msg("DB_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal().Msgf("Cannot connect to database: %s", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Baju{}); err != nil {
		log.Fatal().Msgf("Error during migration: %s", err)
	}

	return db
}
