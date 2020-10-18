package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Constants struct {
	BucketLength uint64
	WindowLength uint64

	PostgresUser string
	PostgresPass string
	PostgresDB   string
	PostgresHost string

	JwtSecret string
	JwtTtl    uint64

	DefaultAdminUsername string
	DefaultAdminPassword string
}

var Consts Constants

func InitConstants(logger *log.Logger) {

	// loading .env file to environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env setting file")
	}

	// set bucket length
	Consts.BucketLength, err = strconv.ParseUint(os.Getenv("BUCKET_LENGTH"), 10, 64)
	// set window length
	Consts.WindowLength, err = strconv.ParseUint(os.Getenv("WINDOW_LENGTH"), 10, 64)
	// set postgres constants
	Consts.PostgresDB = os.Getenv("POSTGRES_DB")
	Consts.PostgresUser = os.Getenv("POSTGRES_USER")
	Consts.PostgresPass = os.Getenv("POSTGRES_PASSWORD")
	Consts.PostgresHost = os.Getenv("POSTGRES_HOST")

	Consts.JwtSecret = os.Getenv("JWT_SECRET")
	Consts.JwtTtl, err = strconv.ParseUint(os.Getenv("JWT_TTL"), 10, 64)

	Consts.DefaultAdminUsername = os.Getenv("DEFAULT_ADMIN_USERNAME")
	Consts.DefaultAdminPassword = os.Getenv("DEFAULT_ADMIN_PASSWORD")

	// log if there was an error converting string to number
	if err != nil {
		logger.Fatal("parsing env file to constant struct failed")
	}

}
