package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInitGorm(t *testing.T) {
	// Load the .env file for the test
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	// Backup the current DB_URL environment variable and restore it after the test
	originalDBURL := os.Getenv("DB_URL")
	defer os.Setenv("DB_URL", originalDBURL)

	// Test case: Valid DB_URL
	os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/test_sagara?sslmode=disable")

	// Run the InitGorm function
	db := InitGorm()

	// Assertions
	assert.NotNil(t, db, "Expected Gorm DB instance to be initialized")
	assert.NoError(t, db.Error, "Expected no error when initializing Gorm DB")

	// Clean up the DB connection after the test
	sqlDB, err := db.DB()
	assert.NoError(t, err, "Expected no error when getting sql.DB from Gorm DB")
	sqlDB.Close()
}
