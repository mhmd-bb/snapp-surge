package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func InitDB(host string, user string, pass string, database string, logger *log.Logger) *gorm.DB {

	var dsn = fmt.Sprintf(
		"user=%s password=%s database=%s host=%s port=5432 sslmode=disable",
		user,
		pass,
		database,
		host)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.WithField("error", err).Fatal("database connection failed")
	}
	logger.Info("connected to database successfully.")
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
