package config

import (
	"os"

	"github.com/anddreluis2/avenida/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	logger = GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat("./db/main.db")
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Account{}, &schemas.Client{})
	if err != nil {
		logger.Errorf("Error migrating schemas: %v", err)
		return nil, err
	}
	return db, nil
}
