package database

import (
	"fmt"

	"github.com/osamikoyo/publics/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBpath))
	if err != nil {
		return nil, fmt.Errorf("cant get db: %v", err)
	}

	return db, nil
}
