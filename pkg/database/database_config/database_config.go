package databaseconfig

import (
	"fmt"

	"github.com/phn00dev/golang-news-app/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	config *config.Config
}

func NewDbConfig(config *config.Config) *DatabaseConfig {
	return &DatabaseConfig{
		config: config,
	}
}

func (dbConfig *DatabaseConfig) GetConfig() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.config.DatabaseConfig.Host,
		dbConfig.config.DatabaseConfig.User,
		dbConfig.config.DatabaseConfig.Password,
		dbConfig.config.DatabaseConfig.Name,
		dbConfig.config.DatabaseConfig.Port,
		dbConfig.config.DatabaseConfig.SslMode,
		dbConfig.config.DatabaseConfig.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
