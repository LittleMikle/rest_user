package storage

import (
	"fmt"
	"github.com/LittleMikle/rest_user/pkg/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

var DB *gorm.DB

func NewConnection(config *Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		config.User, config.Password, config.DBname, config.Host, config.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("failed connect to Postgres: %w", err)
	} else {
		log.Info().Msg("Connection to Postgres successful")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return db, fmt.Errorf("failed automigrate to Postgres: %w", err)
	}

	DB = db
	return DB, nil
}
