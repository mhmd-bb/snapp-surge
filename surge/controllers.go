package surge

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator"
    "net/http"
)

type SurgeController struct {
    surgeService    ISurgeService
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

func (sc *SurgeController)GetAllRules(c *gin.Context) {
    rules, err := sc.surgeService.GetAllRules()

    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "rules fetched successfully", "rules": rules, "status": http.StatusOK})
    return
}

func (sc *SurgeController)CreateRule(c *gin.Context) {

    var ruleDto RuleDto

    err := c.BindJSON(&ruleDto)
    if err != nil {
        _ = c.Error(err)
        return
    }

    var rule Rule

    err = sc.surgeService.CreateRule(&rule, ruleDto)

    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "rule created successfully", "rule": rule, "status": http.StatusCreated})
    return
}

func (sc *SurgeController)DeleteRuleById(c *gin.Context) {

    var deleteRuleDto DeleteRuleDto

    err := c.BindJSON(&deleteRuleDto)
    if err != nil {
        _ = c.Error(err)
        return
    }

    err = sc.surgeService.DeleteRuleById(deleteRuleDto.Id)
    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "rule deleted successfully", "status": http.StatusOK})
    return
}

func NewSurgeController(surgeService ISurgeService) *SurgeController{

    return &SurgeController{surgeService: surgeService}
}