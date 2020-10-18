package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmd-bb/snapp-surge/auth"
	"github.com/mhmd-bb/snapp-surge/config"
	"github.com/mhmd-bb/snapp-surge/database"
	_ "github.com/mhmd-bb/snapp-surge/docs"
	"github.com/mhmd-bb/snapp-surge/logger"
	"github.com/mhmd-bb/snapp-surge/osm"
	"github.com/mhmd-bb/snapp-surge/surge"
	"github.com/mhmd-bb/snapp-surge/user"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, loggerInstance *log.Logger) {

	err := db.AutoMigrate(&surge.Bucket{})
	if err != nil {
		loggerInstance.Error("Bucket migration failed")
	}

	err = db.AutoMigrate(&surge.Rule{})
	if err != nil {
		loggerInstance.Error("Rule migration failed")
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		loggerInstance.Error("User migration failed")
	}

	loggerInstance.Info("Migrations were successful")
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

// @title Surge
// @version 1.0
// @description Snapp Surge Service.

// @host localhost:8080

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {

	// setup logger
	loggerInstance := logger.NewLogger(&log.JSONFormatter{}, "logger/log.txt")

	// initialize config constants
	config.InitConstants(loggerInstance)

	// connect to database
	db := database.InitDB(config.Consts.PostgresUser, config.Consts.PostgresPass, config.Consts.PostgresDB, loggerInstance)

	// migrate all models
	Migrate(db, loggerInstance)

	r := gin.Default()

	// setup auth service
	var jwtAuthService auth.IJwtAuthService = auth.NewJwtService(config.Consts.JwtSecret, config.Consts.JwtTtl, loggerInstance)

	// setup user package
	var usersService user.IUserService = user.NewUsersService(db, jwtAuthService, loggerInstance)
	usersController := user.NewUsersController(usersService)
	usersRouter := user.NewUsersRouter(usersController)

	// create Default admin user if there is no other users in db
	_ = usersService.CreateDefaultUser(config.Consts.DefaultAdminUsername, config.Consts.DefaultAdminPassword)

	// setup osm package
	osmService := osm.NewOpenStreetMapService(db)

	// setup surge package
	var surgeService surge.ISurgeService = surge.NewSurgeService(db, osmService, config.Consts.BucketLength, config.Consts.WindowLength, loggerInstance)
	surgeController := surge.NewSurgeController(surgeService)
	surgeRouter := surge.NewSurgeRouter(surgeController)

	// setup content-type
	r.Use(JSONMiddleware())

	// setup router
	surgeRouter.SetupRouter(r)
	usersRouter.SetupRouter(r)

	// setup swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/swagger.json") // The url pointing to API definition
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Static("/swagger", "docs/")

	r.Run()

}