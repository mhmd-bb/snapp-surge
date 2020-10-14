package surge

import (
    "github.com/gin-gonic/gin"
)

type SurgeRouter struct {
    surgeController      *SurgeController
}

func (sr *SurgeRouter) SetupRouter(r *gin.Engine) *gin.Engine {

    surge := r.Group("/surge")
    {
        surge.POST("ride", sr.surgeController.Ride)
    }

    return r
}

func NewSurgeRouter (surgeController *SurgeController) *SurgeRouter{
    return &SurgeRouter{surgeController: surgeController}
}