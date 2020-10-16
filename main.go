package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mhmd-bb/snapp-surge/auth"
	"github.com/mhmd-bb/snapp-surge/config"
	"github.com/mhmd-bb/snapp-surge/database"
	"github.com/mhmd-bb/snapp-surge/osm"
	"github.com/mhmd-bb/snapp-surge/surge"
	"github.com/mhmd-bb/snapp-surge/user"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&surge.Bucket{})
	db.AutoMigrate(&surge.Rule{})
	db.AutoMigrate(&user.User{})
	fmt.Println("Migrations were successful")
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {

	// initialize config constants
	config.InitConstants()

	// connect to database
	db := database.InitDB(config.Consts.PostgresUser, config.Consts.PostgresPass, config.Consts.PostgresDB)

	// migrate all models
	Migrate(db)

	r := gin.Default()

	// setup auth package
	var jwtAuthService auth.IJwtAuthService = auth.NewJwtService(config.Consts.JwtSecret, config.Consts.JwtTtl)

	// setup user package
	var usersService user.IUserService = user.NewUsersService(db, jwtAuthService)
	usersController := user.NewUsersController(usersService)
	usersRouter := user.NewUsersRouter(usersController)

	// create Default admin user if there is no other users in db
	_ = usersService.CreateDefaultUser(config.Consts.DefaultAdminUsername, config.Consts.DefaultAdminPassword)

	// setup osm package
	osmService := osm.NewOpenStreetMapService(db)

	// setup surge package
	surgeService := surge.NewSurgeService(db, osmService, config.Consts.BucketLength, config.Consts.WindowLength)
	surgeController := surge.NewSurgeController(surgeService)
	surgeRouter := surge.NewSurgeRouter(surgeController)

	// setup content-type
	r.Use(JSONMiddleware())

	// setup router
	surgeRouter.SetupRouter(r)
	usersRouter.SetupRouter(r)

	r.Run()

}