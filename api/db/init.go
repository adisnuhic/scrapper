package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adisnuhic/scrapper_api/config"
	"github.com/jinzhu/gorm"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// scrapper app data store
var scrapperStore Store

// Init initialize db
func Init(cfg *config.AppConfig) {
	env := os.Getenv("ENV")
	scrapperStore = initDB(cfg.DBConnections[env])
}

// Connection get databse connection
func Connection() Store {
	return scrapperStore
}

// Close database connection
func Close() error {
	if scrapperStore != nil {
		return scrapperStore.DB().Close()
	}
	return nil
}

// initDB init database connection
func initDB(dbConn config.DBConnection) Store {
	fmt.Println()
	if dbConn.DBDialect == "" || dbConn.DBConnection == "" {
		return nil
	}

	// open DB connection
	myDB, err := gorm.Open(dbConn.DBDialect, dbConn.DBConnection)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ping database
	if err := myDB.DB().Ping(); err != nil {
		log.Fatal(err.Error())
	}

	// SetMaxIdleConns sets maximum number of connections in the idle connection pool
	maxConn := dbConn.DbMaxIdleConns
	myDB.DB().SetMaxIdleConns(maxConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database
	maxConn = dbConn.DbMaxOpenConns
	myDB.DB().SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxConn = dbConn.DbConnMaxLifetime
	duration := time.Minute * time.Duration(maxConn)
	myDB.DB().SetConnMaxLifetime(duration)

	// Enable Logger, show detailed log
	myDB.LogMode(dbConn.DbLogging)

	log.Println("initialized  API database successfully")

	return myDB

}
