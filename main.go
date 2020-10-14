package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mhmd-bb/snapp-surge/config"
	"github.com/mhmd-bb/snapp-surge/database"
	"github.com/mhmd-bb/snapp-surge/osm"
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

	r := gin.Default()

	// setup osm package
	osmService := osm.NewOpenStreetMapService(db)

	// setup surge package
	surgeService := surge.NewSurgeService(db, osmService, config.Consts.BucketLength)
	surgeController := surge.NewSurgeController(surgeService)
	surgeRouter := surge.NewSurgeRouter(surgeController)

	// setup router
	surgeRouter.SetupRouter(r)

	r.Run()

}