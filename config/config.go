package config

import (
	"gorm.io/gorm"
	"errors"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	// var err error

	// db, err = gorm.Open("sqlite3", "openings.db")
	
	// if err != nil {
	// 	return err
	// }

	return nil 
}

func GetLogger(p string) *Logger {
	// initialize logger
	logger = NewLogger(p)
 
	return logger
}