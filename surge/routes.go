package surge

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmd-bb/snapp-surge/auth"
)

type SurgeRouter struct {
	surgeController *SurgeController
}

func (sr *SurgeRouter) SetupRouter(r *gin.Engine) *gin.Engine {

	surge := r.Group("/surge")
	{
		surge.POST("ride", sr.surgeController.Ride)
	}

	rule := r.Group("/rules", auth.AuthorizeJWT())
	{
		rule.GET("", BadRequestErrorMiddleware(), sr.surgeController.GetAllRules)
		rule.POST("", BadRequestErrorMiddleware(), sr.surgeController.CreateRule)
		rule.DELETE(":id", BadRequestErrorMiddleware(), sr.surgeController.DeleteRuleById)
	}

	return r
}

func NewSurgeRouter(surgeController *SurgeController) *SurgeRouter {
	return &SurgeRouter{surgeController: surgeController}
}
