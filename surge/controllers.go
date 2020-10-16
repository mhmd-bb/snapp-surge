package surge

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator"
    "net/http"
)

type SurgeController struct {
    surgeService    *SurgeService
}

func (sc *SurgeController)Ride(c *gin.Context) {

    // Declare user input Data Transfer Object
    var latLonDto LatLonDto

    // get request body and validate it
    err := c.BindJSON(&latLonDto)

    // return exact error on each field
    if err != nil {

        errors, _ := err.(validator.ValidationErrors)

        e := make(map[string]string)

        for _, err := range errors {
            e[err.Field()] = err.Tag()
        }
        c.JSON(400, e)
        return
    }

    // Get District ID from latitude and longitude
    // if it's not in supported region return appropriate error
    var districtID uint8
    err = sc.surgeService.GetDistrictIDFromLocation(&districtID, latLonDto.Lat, latLonDto.Lon)
    if districtID == 0 {
        c.JSON(http.StatusOK, gin.H{"error": "Latitude and Longitude is not in supported region"})
        return
    }

    // Get Last active bucket of requested district and increment its counter by one
    var lastActiveBucket Bucket
    err = sc.surgeService.IncrementLastActiveBucket(&lastActiveBucket, districtID)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"error": err.Error()})
        return
    }

    // Add all bucket counters in moving window
    var requestsCountInWindow uint64
    err = sc.surgeService.SumAllBucketsInCurrentWindow(&requestsCountInWindow, districtID)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"error": err.Error()})
        return
    }

    // Get Coefficient from count of requests
    var coefficient float32
    err = sc.surgeService.CalculateCoefficient(&coefficient, requestsCountInWindow)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"coefficient": coefficient})
}

func NewSurgeController(surgeService *SurgeService) *SurgeController{

    return &SurgeController{surgeService: surgeService}
}