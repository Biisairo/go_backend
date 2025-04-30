package database

import (
	"clonecoding/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func GetDatabase() *gorm.DB {
	return database
}

func InitDatabase() {
	var err error
	database, err = gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})
	if err != nil {
		panic("Fail to open database")
	}
}

func InitScheme(model any) error {
	return database.AutoMigrate(&model)
}
