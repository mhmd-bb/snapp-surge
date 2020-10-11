package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "strconv"
)

var BucketLength uint64

func Init()  {

    // loading .env file to environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // set bucket length
    BucketLength, _ = strconv.ParseUint(os.Getenv("BUCKET_LENGTH"), 10, 64)


    // init database
    InitDB()


}