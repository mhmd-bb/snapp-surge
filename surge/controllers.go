package surge

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator"
    "net/http"
)

type SurgeController struct {
    surgeService    *SurgeService
    bucketLength    uint64
}

func (sc *SurgeController)Ride(c *gin.Context) {

    var latLonDto LatLonDto

    // get request body and validate it
    err := c.BindJSON(&latLonDto)

    if err != nil {

        errors, _ := err.(validator.ValidationErrors)

        e := make(map[string]string)

        for _, err := range errors {
            e[err.Field()] = err.Tag()
        }
        c.JSON(400, e)
        return
    }

    // get district id from osm service
    districtID := sc.surgeService.GetDistrictIDFromLocation(latLonDto.Lat, latLonDto.Lon)

    var bucket Bucket

    err = sc.surgeService.IncrementLastActiveBucket(&bucket, districtID)

    //TODO sum all buckets in current interval (ex: 2 hours)

    //TODO calculate coefficient

    c.JSON(http.StatusOK, bucket)
}

func NewSurgeController(surgeService *SurgeService, bucketLength uint64) *SurgeController{

    return &SurgeController{surgeService: surgeService, bucketLength: bucketLength}
}