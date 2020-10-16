package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "strconv"
)

type Constants struct {
    BucketLength uint64
    WindowLength uint64

    PostgresUser string
    PostgresPass string
    PostgresDB   string

    JwtSecret string
    JwtTtl    uint64

    DefaultAdminUsername string
    DefaultAdminPassword string
}

var Consts Constants

func InitConstants() {

    // loading .env file to environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // set bucket length
    Consts.BucketLength, _ = strconv.ParseUint(os.Getenv("BUCKET_LENGTH"), 10, 64)
    // set window length
    Consts.WindowLength, _ = strconv.ParseUint(os.Getenv("WINDOW_LENGTH"), 10, 64)
    // set postgres constants
    Consts.PostgresDB = os.Getenv("POSTGRES_DB")
    Consts.PostgresUser = os.Getenv("POSTGRES_USER")
    Consts.PostgresPass = os.Getenv("POSTGRES_PASSWORD")

    Consts.JwtSecret = os.Getenv("JWT_SECRET")
    Consts.JwtTtl, _ = strconv.ParseUint(os.Getenv("JWT_TTL"), 10, 64)

    Consts.DefaultAdminUsername = os.Getenv("DEFAULT_ADMIN_USERNAME")
    Consts.DefaultAdminPassword = os.Getenv("DEFAULT_ADMIN_PASSWORD")

}
