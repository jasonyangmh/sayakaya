package database

import (
	"fmt"
	"log"

	"github.com/jasonyangmh/sayakaya/config"
	"github.com/jasonyangmh/sayakaya/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Database *gorm.DB
}

func Connect(cfg *config.Config) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
		return nil
	}

	return &Database{
		Database: db,
	}
}

func (db *Database) Migrate() {
	db.Database.Migrator().DropTable(
		&model.User{},
		&model.Promo{},
		&model.UserPromo{},
	)

	db.Database.AutoMigrate(
		&model.User{},
		&model.Promo{},
		&model.UserPromo{},
	)
}
