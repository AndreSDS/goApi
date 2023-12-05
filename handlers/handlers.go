package handlers

import (
	"github.com/AndreSDS/goApi/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handlers")
	db = config.GetSQLite()
}