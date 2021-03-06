package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

var (
	apiConfig *Config
	doOnce    sync.Once
)

// AppServerConfig - configs of app server
type AppServerConfig struct {
	Host string `envconfig:"HOST" default:"localhost:"`
	Port string `envconfig:"PORT" default:"9090"`
	SwaggerHost string `envconfig:"SWAGGER_HOST" default:"localhost:9090"`
}

// DatabasePostgresConfig - configs of database
type DatabasePostgresConfig struct {
	DatabaseURL  string `envconfig:"DATABASE_URL"`
	Host         string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port         string `envconfig:"DATABASE_PORT" default:"5439"`
	UserName     string `envconfig:"DATABASE_USER" default:"postgres"`
	Password     string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	DatabaseName string `envconfig:"DATABASE_NAME" default:"digital_accounts"`
	SSLMode      string `envconfig:"DATABASE_SSLMODE" default:"disable"`
	PoolMinSize  string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize  string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
}

// AuthConfig - configs of authentication
type AuthConfig struct {
	JWTSecretKey     string `envconfig:"JWT_SECRET_KEY" default:"jwt-digital-accounts-secret-key"`
	MaxTokenLiveTime string `envconfig:"JWT_MAX_TOKEN_LIVE_TIME" default:"50m"` // 50 minutes in Duration format
}

// Config - configs of api
type Config struct {
	Profile          string `envconfig:"PROFILE" default:"dev"`
	ExecuteMigration bool `envconfig:"EXECUTE_MIGRATION" default:"true"`
	MigrationPath    string `envconfig:"MIGRATION_PATH" default:"app/persistence/database/postgres/migrations"`
	AuthConfig       AuthConfig
	WebServerConfig  AppServerConfig
	DatabaseConfig   DatabasePostgresConfig
}

// GetAPIConfigs loads and get environment variables to configure the api
func GetAPIConfigs() *Config {
	var config Config

	doOnce.Do(func() {
		err := envconfig.Process("", &config)

		if err != nil {
			log.Fatalln("Unable to load api configuration")
		}
		apiConfig = &config
	})
	return apiConfig
}

// GetWebServerAddress - returns the web server address
func (webServerConfig AppServerConfig) GetWebServerAddress() string {
	return fmt.Sprintf(
		"%s:%s",
		webServerConfig.Host,
		webServerConfig.Port,
	)
}

// GetDatabaseDSN - returns the database dsn
func (databaseConfig DatabasePostgresConfig) GetDatabaseDSN() string {

	// if DatabaseURL is set, then returns the complete url
	if databaseConfig.DatabaseURL != "" {
		return databaseConfig.DatabaseURL
	}

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

// GetDatabaseURI - returns the database uri
func (databaseConfig DatabasePostgresConfig) GetDatabaseURI() string {

	// if DatabaseURL is set, then returns the complete url
	if databaseConfig.DatabaseURL != "" {
		return databaseConfig.DatabaseURL
	}

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
