package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitDatabase() error {
	return nil
}