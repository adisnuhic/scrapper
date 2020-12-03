package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

// DBConnection connection to a database
type DBConnection struct {
	DBDialect         string
	DBConnection      string
	DbMaxIdleConns    int
	DbMaxOpenConns    int
	DbConnMaxLifetime int
	DbLogging         bool
}

// AppConfig application configuration
type AppConfig struct {
	DBConnections map[string]DBConnection
}

// Load app configuration
func Load() *AppConfig {
	gotenv.Load()

	return &AppConfig{
		DBConnections: map[string]DBConnection{
			"development": {
				DBDialect:         "mysql",
				DBConnection:      os.Getenv("DB_DEV_CONNECTION"),
				DbMaxIdleConns:    10,
				DbMaxOpenConns:    100,
				DbConnMaxLifetime: 30, // minutes
				DbLogging:         true,
			},
		},
	}
}

// InitLogger configuration for logger
func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
}
