package main

import (
	"fmt"
	"github.com/mhmd-bb/snapp-surge/config"
	"github.com/mhmd-bb/snapp-surge/surge"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&surge.Bucket{})

	fmt.Println("Migrations were successful")
}

func main() {

	// initialize config
	config.Init()

	db:=config.GetDB()
	Migrate(db)

}