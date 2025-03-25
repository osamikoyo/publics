package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		return nil, fmt.Errorf("cant get db: %v", err)
	}

	return db, nil
}
