package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func InitDB(user string, pass string, database string) *gorm.DB {

	var dsn = fmt.Sprintf(
		"user=%s password=%s database=%s host=localhost port=5432 sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database err: ", err)
	}
	fmt.Println("connected to database successfully.")
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}