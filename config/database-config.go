package config

import (
	"fmt"
	"os"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5434 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to create database connection")
	}

	db.AutoMigrate(
		&entity.RegisterVerification{},
		&entity.User{},
		&entity.Customer{},
		&entity.History{},
	)
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
