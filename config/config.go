package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "strconv"
)

type Constants struct {
    BucketLength    uint64

    PostgresUser    string
    PostgresPass    string
    PostgresDB    string

}

var Consts Constants

func InitConstants()  {

    // loading .env file to environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // set bucket length
    Consts.BucketLength, _ = strconv.ParseUint(os.Getenv("BUCKET_LENGTH"), 10, 64)

}