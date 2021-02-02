package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var AppConfig *Config

// WebServerConfig - configs of web server
type WebServerConfig struct {
	Host string `envconfig:"WEB_SERVER_HOST" default:"localhost"`
	Port string `envconfig:"WEB_SERVER_PORT" default:"9090"`
}

// DatabasePostgresConfig - configs of database
type DatabasePostgresConfig struct {
	Host         string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port         string `envconfig:"DATABASE_PORT" default:"5434"`
	UserName     string `envconfig:"DATABASE_USER" default:"postgres"`
	Password     string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	DatabaseName string `envconfig:"DATABASE_NAME" default:"digital_accounts"`
	SSLMode      string `envconfig:"DATABASE_SSLMODE" default:"disable"`
	PoolMinSize  string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize  string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
}

// AuthConfig - configs of authentication
type AuthConfig struct {
	JwtSecretKey     string `envconfig:"JWT_SECRET_KEY" default:"jwt-digital-accounts-secret-key"`
	MaxTokenLiveTime string `envconfig:"JWT_MAX_TOKEN_LIVE_TIME" default:"10"` // Minutes
}

// Config - configs of api
type Config struct {
	Profile         string `envconfig:"PROFILE" default:"dev"`
	AuthConfig      AuthConfig
	WebServerConfig WebServerConfig
	DatabaseConfig  DatabasePostgresConfig
}

// LoadConfigs loads environment variables to configure the api
func LoadConfigs() *Config {
	var config Config
	err := envconfig.Process("", &config)

	if err != nil {
		log.Fatalln("Unable to load api configuration")
	}
	AppConfig = &config
	return AppConfig
}

// GetWebServerAddress - returns the web server address
func (webServerConfig WebServerConfig) GetWebServerAddress() string {
	return fmt.Sprintf(
		"%s:%s",
		webServerConfig.Host,
		webServerConfig.Port,
	)
}

// GetDatabaseDSN - returns the database dsn
func (databaseConfig DatabasePostgresConfig) GetDatabaseDSN() string {
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

// GetDatabaseURL - returns the database url
func (databaseConfig DatabasePostgresConfig) GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		databaseConfig.UserName,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.DatabaseName,
		databaseConfig.SSLMode,
	)
}
