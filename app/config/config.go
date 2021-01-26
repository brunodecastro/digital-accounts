package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type WebServerConfig struct {
	Host string `envconfig:"WEB_SERVER_HOST" default:"localhost"`
	Port string `envconfig:"WEB_SERVER_PORT" default:"9090"`
}

type DatabaseConfig struct {
	Host         string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port         string `envconfig:"DATABASE_PORT" default:"5434"`
	UserName     string `envconfig:"DATABASE_USER" default:"postgres"`
	Password     string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	DatabaseName string `envconfig:"DATABASE_NAME" default:"digital_accounts"`
	SSLMode      string `envconfig:"DATABASE_SSLMODE" default:"disable"`
	PoolMinSize  string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize  string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
}

type Config struct {
	Profile         string `envconfig:"PROFILE" default:"dev"`
	WebServerConfig WebServerConfig
	DatabaseConfig  DatabaseConfig
}

// LoadConfigs loads environment variables to configure the api
func LoadConfigs() *Config {
	var config Config
	err := envconfig.Process("", &config)

	if err != nil {
		log.Fatalln("Unable to load api configuration")
	}
	return &config
}

func (webServerConfig WebServerConfig) GetWebServerAddress() string {
	return fmt.Sprintf(
		"%s:%s",
		webServerConfig.Host,
		webServerConfig.Port,
	)
}

func (databaseConfig DatabaseConfig) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%s pool_max_conns=%s",
		databaseConfig.UserName,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.DatabaseName,
		databaseConfig.SSLMode,
		databaseConfig.PoolMinSize,
		databaseConfig.PoolMaxSize,
	)
}
