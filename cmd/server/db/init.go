package db

import (
	"fmt"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB(cfg config.DatabaseConfiguration) (*gorm.DB, error) {
	switch cfg.Dialect {
	case "sqlite":
		db, err := initSqlite(cfg)
		if err != nil {
			return nil, fmt.Errorf("error initializing sqlite: %v", err)
		}

		return db, nil
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", cfg.Dialect)
	}
}

func initSqlite(cfg config.DatabaseConfiguration) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Datasource), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if cfg.AutoMigrate {
		err = db.AutoMigrate(&models.MessageCount{})
		if err != nil {
			return nil, fmt.Errorf("error migrating database: %v", err)
		}
	}

	return db, nil
}
