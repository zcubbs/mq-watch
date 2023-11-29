package db

import (
	"fmt"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
	"github.com/zcubbs/mq-watch/cmd/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	log = logger.L()
)

func InitializeDB(cfg config.DatabaseConfiguration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	if cfg.Sqlite.Enabled && cfg.Postgres.Enabled {
		return nil, fmt.Errorf("both sqlite and postgres are enabled, please only enable one")
	}

	if cfg.Sqlite.Enabled {
		log.Info("Initializing SQLite database")
		db, err = initSqlite(cfg.Sqlite)
	} else if cfg.Postgres.Enabled {
		log.Info("Initializing PostgreSQL database")
		db, err = initPostgres(cfg.Postgres)
	} else {
		return nil, fmt.Errorf("no database enabled")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initSqlite(cfg config.SQLiteConfiguration) (*gorm.DB, error) {
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

func initPostgres(cfg config.PostgresConfiguration) (*gorm.DB, error) {
	// PostgreSQL connection string format: "host=myhost user=myuser password=mypass dbname=mydb sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening postgres database: %v", err)
	}

	// Perform auto-migration if needed
	if cfg.AutoMigrate {
		err = db.AutoMigrate(&models.MessageCount{})
		if err != nil {
			return nil, fmt.Errorf("error migrating postgres database: %v", err)
		}
	}

	return db, nil
}
