package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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

// GRPCConnection connection to grpc server
type GRPCConnection struct {
	GRPCOptions []grpc.DialOption
	Address     string
}

// AppConfig application configuration
type AppConfig struct {
	DBConnections   map[string]DBConnection
	GRPCConnections map[string]GRPCConnection
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
		GRPCConnections: map[string]GRPCConnection{
			"development": {
				GRPCOptions: []grpc.DialOption{grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
					Time:                20 * time.Second,
					PermitWithoutStream: true,
				})},
				Address: os.Getenv("DEV_GRPC_SERVER"),
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
