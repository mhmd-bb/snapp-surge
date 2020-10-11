package main

import (
	"fmt"
	"github.com/mhmd-bb/snapp-surge/config"
	"github.com/mhmd-bb/snapp-surge/database"
	"github.com/mhmd-bb/snapp-surge/surge"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&surge.Bucket{})

	fmt.Println("Migrations were successful")
}

func main() {

	// initialize config constants
	config.InitConstants()

	// connect to database
	db := database.InitDB(config.Consts.PostgresUser, config.Consts.PostgresPass, config.Consts.PostgresDB)

	// migrate all models
	Migrate(db)


	db.Create(&surge.Bucket{BucketLength: 600})

}