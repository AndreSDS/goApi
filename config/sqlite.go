package config

import (
	"github.com/AndreSDS/goApi/schemas"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it...")

		// create the database file and directory
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	// create the database and connect to it
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to migrate database: %v", err)
		return nil, err	
	}

	// return the database
	return db, nil
}