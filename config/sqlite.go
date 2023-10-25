package config

import (
	"os"

	"github.com/ramown/gopportunites/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitalizeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the dabase file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the database file and directory
		if err = os.MkdirAll("./db", os.ModePerm); err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create database connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening erro: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration erro: %v", err)
		return nil, err
	}

	return db, nil
}
