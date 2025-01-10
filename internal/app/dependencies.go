package app

import (
	"net/http"

	"github.com/phn00dev/golang-news-app/pkg/config"
	databaseconfig "github.com/phn00dev/golang-news-app/pkg/database/database_config"
	httpclient "github.com/phn00dev/golang-news-app/pkg/httpClient"
	redisconfig "github.com/phn00dev/golang-news-app/pkg/redisConfig"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB          *gorm.DB
	HttpClient  *http.Client
	Config      *config.Config
	RedisConfig *redisconfig.RedisConfig
}

func GetDependencies() (*Dependencies, error) {
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	// postgres connection
	newPostgresDB := databaseconfig.NewDbConfig(getConfig)
	getDB, err := newPostgresDB.GetConfig()
	if err != nil {
		return nil, err
	}

	// redis connection

	// http client connection
	clientHttp := httpclient.NewHttpConnect()

	return &Dependencies{
		DB:         getDB,
		HttpClient: clientHttp,
		Config:     getConfig,
	}, nil
}
