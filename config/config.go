package config

import (
	"gorm.io/gorm"
	"fmt"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// initialize SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("failed to initialize SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	// return SQLite database
	return db
}

func GetLogger(p string) *Logger {
	// initialize logger
	logger = NewLogger(p)
 
	return logger
}